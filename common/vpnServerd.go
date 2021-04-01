package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var LogOut chan string
var LoginChan chan string

var vpnipbindDomain map[string]string

type Gcurrlife struct {
	GM map[string]string
	sync.RWMutex
}

var currlifeuser *Gcurrlife
var GDB *gorm.DB

// var currlifeuser = InitGcurrlife()
// vpn的账号

type VpnAccountInfo struct {
	gorm.Model

	Username string `gorm:"unique;not null;size:255"`   //vpn 账号
	Passwd   string `gorm:"type:varchar(100);not null"` //密码 // 显示则转换成域名组名
	Telnum   string `gorm:"type:varchar(100)"`          //手机号码
	DisIp    string `gorm:"type:varchar(100);not null"` //分配的ip
	// Email     string     `gorm:"type:varchar(100);unique_index"`
	Stats       string `gorm:"type:varchar(100)"` //是否在登录//
	CreateTimes string `gorm:"type:varchar(100)"` // 账号创建时间
	Lastlogins  string `gorm:"type:varchar(100)"` //上次登录时间
	LoginCount  uint   //总共登录的次数
	LoginPlat   string `gorm:"type:varchar(100)"` //本次登录到的平台
	LoginDura   string `gorm:"type:varchar(100)"` //本次的登录时长
	LoginOut    string `gorm:"type:varchar(100)"` //登出时间

	//外键
	// DomainPermissions []DomainPermission `gorm:"ForeignKey:VpnAcDoid"`
	VpnAcDoid   uint
	VpnAcDoname string `gorm:"type:varchar(100)"`
	NoticeInit  string `gorm:"type:varchar(10)"`
	// PermissionDisgroup []PermissionDisgroup `gorm:"foreignkey:VpnAcDoid;association_foreignkey:VpnAcDoid"`
	// DomainPermissionID int
}

// 域名与属组权限控制

type DomainPermission struct {
	gorm.Model
	Perdomk string
	Domain  string `gorm:"type:varchar(100);unique;not null"`
}

//权限分配组
type PermissionDisgroup struct {
	gorm.Model
	GroupName      string           `gorm:"type:varchar(150);unique;not null"` // 组名
	VpnAcDoid      uint             //
	Perdomk        string           // 域名的主键Perdomkid[1,2,3,4]
	VpnAccountInfo []VpnAccountInfo `gorm:"foreignkey:VpnAcDoid;association_foreignkey:VpnAcDoid"`
}

func init() {

	//gorm
	gdb, err := gorm.Open("mysql", "root:123123@(1.1.1.1:3306)/VpnAudit?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm conn is fail...", err)
		panic(err)
	}

	gdb.SingularTable(true)
	gdb.AutoMigrate(&VpnAccountInfo{}, &PermissionDisgroup{}, &DomainPermission{})
	GDB = gdb
	fmt.Println("初始化数据库 gorm 成功")
	//gorm
}

func InitGcurrlife() *Gcurrlife {
	return &Gcurrlife{GM: make(map[string]string, 50)}
}

func (this *Gcurrlife) Get(username string) bool {
	this.RLock()
	defer this.RUnlock()
	_, exists := this.GM[username]
	return exists
}

func (this *Gcurrlife) Del(username string) {
	this.Lock()
	defer this.Unlock()
	delete(this.GM, username)
}

func (this *Gcurrlife) Put(username, virtualip string) {
	this.Lock()
	defer this.Unlock()
	this.GM[username] = virtualip //存在
}

func ReadLine(fileName string) (error, map[string]string) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err, nil
	}
	buf := bufio.NewReader(f)
	Realcon := false
	var currFilecon map[string]string
	currFilecon = make(map[string]string, 50)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.Contains(line, "GLOBAL STATS") {
			Realcon = false
		}
		if Realcon {
			conui := strings.Split(line, ",")
			currFilecon[conui[1]] = conui[2]
			LoginChan <- line
		}
		if strings.Contains(line, "Virtual Address,Common Name,Real Address,Last Ref") {
			Realcon = true
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err, nil
		}
	}
	return nil, currFilecon
}

func ReadFile(filename string) error {
	t := time.NewTicker(time.Second * time.Duration(2))
	defer t.Stop()
	// numline := 0
	for {
		<-t.C
		err, _ := ReadLine(filename)
		if err != nil {
			fmt.Println(err)
			return err
		}
		// numline++
	}
}

//域名解析
func domainAnaly(domain string) (error, []string) {
	timeout := time.Duration(4 * time.Second)
	done := make(chan error, 1)
	curip := make([]string, 0, 5)
	go func() {
		ip, err := net.LookupHost(domain)
		curip = ip
		done <- err
	}()

	select {
	case <-time.After(timeout):
		fmt.Println(domain, "域名解析超时5s")
		return errors.New("domain analysis is timeout"), curip
	case err := <-done:
		if err != nil {
			fmt.Println(domain, "解析失败!")
			return err, curip
		}
	}
	return nil, curip

}

func SearchUserBindIP(username string) (error, []string) {
	var va VpnAccountInfo
	var pd PermissionDisgroup
	var currUserAllIp []string
	err := GDB.Where("username = ?", username).First(&va).Error
	if err != nil {
		fmt.Println("查询用户失败!", err)
		return err, nil
	}
	err = GDB.Where("vpn_ac_doid = ?", va.VpnAcDoid).First(&pd).Error
	if err != nil {
		fmt.Println("查询用户对应的域名组失败:", err)

		return err, nil
	}
	for _, val := range strings.Split(pd.Perdomk, ",") {
		// fmt.Println("val", val)
		var dp DomainPermission
		err = GDB.Where("perdomk = ?", val).First(&dp).Error
		if err != nil {
			fmt.Println("查询用户对应的域名失败:", err)
			return err, nil
		}
		fmt.Println("当前用户的对应的域名有:", dp.Domain)
		err, doanaip := domainAnaly(dp.Domain)
		if err != nil {
			fmt.Println("域名解析失败:", err)
			return err, nil
		}
		currUserAllIp = append(currUserAllIp, doanaip...)

	}
	return nil, currUserAllIp
}

func CreateIptablesRuls(res []string) {
	fmt.Println(res, "创建防火墙规则")
	err, currips := SearchUserBindIP(res[1])
	if err != nil {
		fmt.Println("获取域名对应的ip失败:", err)
		return
	}
	// fmt.Println("currips....", currips)
	for _, val := range currips {
		cmdIptables := fmt.Sprintf("iptables -I FORWARD -s %s/32 -d %s/32 -j DROP", res[2], val)
		fmt.Println("cmdIptables", cmdIptables)
	}

	// cmd := exec.Command("iptables", fmt.Sprintf("-I FORWARD -s %s/32 -d %s/32 -j DROP", res[1], "1.1.1.1"))
	// var stdout bytes.Buffer
	// cmd.Stdout = &stdout
	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("使用RPC 通知")
	// }

}

func CleanIptablesRuls(res interface{}) {
	matchIptablesip := res.(string)
	cmd := exec.Command("iptables", "-L")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("防火墙命令失败  iptables -l", err)
	}
	fmt.Println(stdout.String(), "这是当前防火墙的规则")
	fmt.Println("循环读取防火墙规则,删除有matchIptablesip", matchIptablesip)
	// 循环读取防火墙规则,删除 有matchIptablesip 的内容
	// fmt.Println("15s 清除防火墙规则,当前内存的数据为:") //res.(map[string]string)

	// cmd := exec.Command("iptables", fmt.Sprintf("-I FORWARD -s %s/32 -d %s/32 -j DROP", res[1], "1.1.1.1"))
	// var stdout bytes.Buffer
	// cmd.Stdout = &stdout
	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("使用RPC 通知")
	// }

}

func fromLogincon() {
	for {
		select {
		case line := <-LoginChan:
			conui := strings.Split(line, ",")
			exists := currlifeuser.Get(conui[1])
			if !exists {
				go CreateIptablesRuls(conui)
			}
			currlifeuser.Put(conui[1], conui[2])
		}
	}
}

func fromLogOutcon() {
	for {
		select {
		case res := <-LogOut:
			CleanIptablesRuls(res)
		}
	}
}

func CleanLoginmsg() {
	time.Sleep(time.Second * 10)
	for {

		err, currFilecon := ReadLine("openvpn.log")
		if err != nil {
			fmt.Println(err)
			continue //后面再处理,启动服务调用
		}
		for k, v := range currlifeuser.GM {
			_, exist := currFilecon[k]
			if !exist {
				currlifeuser.Del(k)
				CleanIptablesRuls(v)
			}
		}

		fmt.Println("当前内存的加载的信息:", currlifeuser.GM)

		time.Sleep(time.Second * 15)

	}
}

// //超时处理
// func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
// 	var err error
// 	done := make(chan error)
// 	go func() {
// 		done <- cmd.Wait()
// 	}()
// 	select {
// 	case <-time.After(timeout):
// 		log.Printf("timeout, process:%s will be killed", cmd.Path)

// 		go func() {
// 			<-done // allow goroutine to exit
// 		}()
// 		//IMPORTANT: cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} is necessary before cmd.Start()
// 		err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
// 		if err != nil {
// 			log.Println("kill failed, error:", err)
// 		}
// 		return err, true
// 	case err = <-done:
// 		return err, false
// 	}
// }

// func ReloadMsg() {
// 	for {
// 		time.Sleep(time.Second * 30)
// 		fmt.Println("重新加载用户域名对应的信息!")
// for k, _ := range currlifeuser.GM {
// 	SearchUserBindIP(username)
// }

// 	}
// }
func main() {
	currlifeuser = InitGcurrlife()
	go ReadFile("openvpn.log")
	// vpnipbindDomain = make(map[string]string, 100)

	LoginChan = make(chan string, 100)
	// LogOut = make(chan string, 100)
	go fromLogincon()
	go CleanLoginmsg()
	// go fromLogOutcon()
	// go ReloadMsg() //30s 重新刷新一次,防止用户对应的域名有所更改 // 后面更改为让使用者断开1分钟重新再连接

	select {}
	// r := gin.Default()
	// r.GET("/userdomainschange", func(c *gin.Context) {
	// 	cdomains := c.Query("cdomains")

	// 	c.JSON(200, gin.H{
	// 		"code": 200,
	// 	})
	// })
	// r.Run()
}
