package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hpcloud/tail"
	"github.com/jinzhu/gorm"
	// "os/exec"
)

var conLoginChan chan string // iptables 允许
// var conLogoutChan chan string // iptables 不允许

var DB *sql.DB

var GDB *gorm.DB

var idmapdomains map[string]string

var vpnAccbindip map[string]string //、登陆的vpn account 和 源ip 保存 在 内存中

type AccountInfo struct {
	gorm.Model
	Username string `gorm:"unique;not null;size:255"`
	Passwd   string `gorm:"type:varchar(100)"`
	// Email     string     `gorm:"type:varchar(100);unique_index"`
	Stats       string `gorm:"type:varchar(100)"`
	CreateTimes string `gorm:"type:varchar(100)"`
	Lastlogins  string `gorm:"type:varchar(100)"`
}

// vpn的账号

type VpnAccountInfo struct {
	gorm.Model

	Username string `gorm:"unique;not null;size:255"`   //vpn 账号
	Passwd   string `gorm:"type:varchar(100);not null"` //密码 // 显示则转换成域名组名

	DisIp string `gorm:"type:varchar(100);not null"` //分配的ip
	// Email     string     `gorm:"type:varchar(100);unique_index"`
	Stats       string `gorm:"type:varchar(100)"` //是否在登录
	CreateTimes string `gorm:"type:varchar(100)"` // 账号创建时间
	Lastlogins  string `gorm:"type:varchar(100)"` //上次登录时间
	LoginCount  uint   //总共登录的次数
	LoginPlat   string `gorm:"type:varchar(100)"` //本次登录到的平台
	LoginDura   string `gorm:"type:varchar(100)"` //本次的登录时长
	LoginOut    string `gorm:"type:varchar(100)"` //登出时间

	//外键
	// DomainPermissions []DomainPermission `gorm:"ForeignKey:VpnAcDoid"`
	VpnAcDoid uint
	// PermissionDisgroup []PermissionDisgroup `gorm:"foreignkey:VpnAcDoid;association_foreignkey:VpnAcDoid"`
	// DomainPermissionID int
}

// 域名与属组权限控制

type DomainPermission struct {
	gorm.Model
	Perdomk string
	Domain  string `gorm:"type:varchar(100);not null"`
}

//权限分配组
type PermissionDisgroup struct {
	gorm.Model
	GroupName      string           `gorm:"type:varchar(150);not null"` // 组名
	VpnAcDoid      uint             //
	Perdomk        string           // 域名的主键Perdomkid[1,2,3,4]
	VpnAccountInfo []VpnAccountInfo `gorm:"foreignkey:VpnAcDoid;association_foreignkey:VpnAcDoid"`
}

//账号数据库
func init() {

	//gorm
	gdb, err := gorm.Open("mysql", "root:123123@(1.2.3.4.5:22111)/vpn-system?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm conn is fail...")
		panic(err)
	}
	gdb.SingularTable(true)
	gdb.AutoMigrate(&AccountInfo{}, &VpnAccountInfo{}, &DomainPermission{}, &PermissionDisgroup{})

	GDB = gdb
	//gorm
}

func CreateIptablesRuls(info string) {
	// vpnSourceIp, sigdomain
	vpnSourcecipbindsigd := strings.Split(info, ",")
	// fmt.Println("管道里的数据")
	// fmt.Println(vpnSourcecipbindsigd)
	iptablesCmdLogin := fmt.Sprintf("iptables -I FORWARD -s %s/32 -d %s/32 -j DROP", vpnSourcecipbindsigd[0], vpnSourcecipbindsigd[1])

	// cmd := exec.Command("/bin/bash", "-c",iptablesCmd)

	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("Error:The command is err,", err)
	// 	return  err
	// }
	// return nil

	fmt.Println("最后的防火墙规则:....", iptablesCmdLogin)

}

//登出

func fromChanLogoutcon() {

	for range time.Tick(5 * time.Second) {

		//每隔五秒执行一次清除防火墙任务

		// fmt.Println("读取整个文件")
		var currFileVpnUser map[string]string
		currFileVpnUser = make(map[string]string, 50)
		f, err := os.Open("openvpn.log")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')

			if err != nil || io.EOF == err {
				break
			}
			// fmt.Println(line)
			sigsliec := strings.Split(line, ",")
			// fmt.Println(sigsliec)
			// fmt.Printf("%T,%v", sigsliec, sigsliec)
			// fmt.Println(len(sigsliec))
			// time.Sleep(time.Second * 1)
			// currFileVpnUser["liuzhuang"] = "1.1.1.1"
			// fmt.Println("读取文件内容:", sigsliec, sigsliec[0], sigsliec[1])
			if len(sigsliec) > 2 {
				currFileVpnUser[sigsliec[1]] = sigsliec[0]
			}

		}

		fmt.Println("内存中的信息....", vpnAccbindip)
		fmt.Println("文件中的信息....", currFileVpnUser)
		for username := range vpnAccbindip {
			sourceip, ok := currFileVpnUser[username]
			if !ok {
				fmt.Println("文件里没有内存中的账号!", username)

				var vi VpnAccountInfo
				var pd PermissionDisgroup

				GDB.Debug().Where("username = ?", username).First(&vi)
				if (VpnAccountInfo{} == vi || vi.VpnAcDoid == 0) {
					fmt.Println("组的结构体为空,没有查到!")
					return

				}

				GDB.Debug().Where("vpn_ac_doid = ?", vi.VpnAcDoid).First(&pd)

				pergrosli := strings.Split(pd.Perdomk, ",")

				for _, item := range pergrosli {

					sigdomain := idmapdomains[item] // id  解析的ip
					// chandata := vpnSourceIp + "," + sigdomain  // 源ip 对应解析的ip
					iptablesCmdLogout := fmt.Sprintf("iptables -D FORWARD -s %s/32 -d %s/32 -j DROP", sourceip, sigdomain)
					fmt.Println("登出防护墙规则....", iptablesCmdLogout)
				}

				//删除自己的表

				delete(vpnAccbindip, username)
				fmt.Println("删除自己的表..........", username, vpnAccbindip)

			}

		}

	}

}

//登陆
func fromChanLogincon() {
	fmt.Println("conLoginChan...登录管道")
	for {
		select {
		case res := <-conLoginChan:
			// 去生成防火墙规则,并执行
			fmt.Println("管道里获取数后放到CreateIptablesRuls函数中", res)
			CreateIptablesRuls(res)
			// if err != nil {
			// 	// alert("报警")
			// 	fmt.Println("报警")
			// }
		default:
			time.Sleep(time.Millisecond * 100)
			// fmt.Println("chan 睡眠200 微妙")
		}
	}
}

func RuleSplic(vpnAccount, vpnSourceIp string) {
	// fmt.Println("从数据库里面查vpn名称,再把组查出来，然后拼接成iptables规则")
	// fmt.Println(vpnAccount, vpnSourceIp)
	// gorm.

	var vi VpnAccountInfo
	var pd PermissionDisgroup
	// var dp DomainPermission
	GDB.Debug().Where("username = ?", vpnAccount).First(&vi)
	if (VpnAccountInfo{} == vi || vi.VpnAcDoid == 0) {
		fmt.Println("组的结构体为空,没有查到!")
		return

	}
	// fmt.Println("账号查询:", vpnAccount, " ", "外键组:", vi.VpnAcDoid)
	GDB.Debug().Where("vpn_ac_doid = ?", vi.VpnAcDoid).First(&pd)
	// if (pd == PermissionDisgroup{}) {
	// 	fmt.Println("域名的结构体为空,没有查到!")

	// }
	// fmt.Println("权限组:", pd.Perdomk)
	pergrosli := strings.Split(pd.Perdomk, ",")
	// fmt.Printf("%T,%v", pergrosli, pergrosli)

	// fmt.Println("解析完后的结果idmapdomains：", idmapdomains)
	for _, item := range pergrosli {

		sigdomain := idmapdomains[item]           // id  解析的ip
		chandata := vpnSourceIp + "," + sigdomain // 源ip 对应解析的ip
		// fmt.Println("将要放到管道里", chandata)
		// if tag == "login"{

		// }
		conLoginChan <- chandata
	}

	// GDB.Debug().Where("vpn_ac_doid = ?", pd.VpnAcDoid).First(&dp)

}

func GetDomainMpas() {
	//域名映射
	idmapdomains = make(map[string]string, 10)
	var dp []DomainPermission
	GDB.Find(&dp)

	for _, item := range dp {
		// idmapdomains[item.Perdomk] = item.Domain //dig获取域名
		addr, err := net.ResolveIPAddr("ip", item.Domain)
		if err != nil {
			fmt.Println("域名解析错误", err.Error())

		}
		// fmt.Println("域名解析:", item.Domain, " : ", addr.String())

		idmapdomains[item.Perdomk] = addr.String()
	}

	fmt.Println("所有对应的ip解析:", idmapdomains)

}

func main() {

	vpnAccbindip = make(map[string]string, 50)
	// go GetDomainMpas()  //先去解析

	GetDomainMpas()
	go fromChanLogincon()

	go fromChanLogoutcon()

	defer GDB.Close()

	conLoginChan = make(chan string, 10)

	t, err := tail.TailFile("openvpn.log", tail.Config{Follow: true, Location: &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}})

	if err != nil {
		fmt.Println("文件读取失败:", err)
	}

	for {
		line, ok := <-t.Lines
		if !ok {
			fmt.Println("本次读取失败!")
			continue
		}

		// conChan <- line.Text
		// 去执行数据库查询防火墙规则拼接
		con := line.Text
		// fmt.Println("con....", con)
		vpnSourceIp := strings.Split(con, ",")[0]
		vpnAccount := strings.Split(con, ",")[1]

		vpnAccbindip[vpnAccount] = vpnSourceIp

		go RuleSplic(vpnAccount, vpnSourceIp)

	}

}
