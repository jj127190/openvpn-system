package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "net/http"
    "openvpn-system/common/share"
    "openvpn-system/dao"
    "strconv"
    "time"
)

//添加账号信息

type ResMsg struct {
    Code int    `json:"code"`
    Mess string `json:"msg"`
}

func AdduserinfoHandler(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获异常:", err)
            msg := ResMsg{
                Code: 500,
                Mess: fmt.Sprintf("%v", err),
            }
            c.JSON(http.StatusOK, msg)
        }
    }()

    username := c.DefaultPostForm("username", "null")
    passwd := c.DefaultPostForm("passwd", "null")
    // fmt.Println(username, passwd)
    share.Logger.Warn("账号添加]..." + fmt.Sprintf("username: %v, passwd: %v", username, passwd))
    if username != "null" && passwd != "null" {

        userinfos := dao.AccountInfo{
            Username:    username,
            Passwd:      passwd,
            Stats:       "created",
            CreateTimes: time.Now().Format("2006-01-02 15:04:05"),
            Lastlogins:  time.Now().Format("2006-01-02 15:04:05"),
        }
        dao.GDB.Create(&userinfos)
    }

    msg := ResMsg{
        Code: 200,
        Mess: "return ok",
    }
    c.JSON(http.StatusOK, msg)

}

//删除账号
func DeluserinfoHandler(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获异常:", err)
            msg := ResMsg{
                Code: 500,
                Mess: fmt.Sprintf("%v", err),
            }
            c.JSON(http.StatusOK, msg)
        }
    }()

    msg := ResMsg{
        Code: 200,
        Mess: "return ok",
    }
    username := c.DefaultQuery("username", "null")
    fmt.Println("删除账号")
    share.Logger.Warn("删除账号]..." + fmt.Sprintf("username: %v", username))
    if username != "null" {

        dao.GDB.Debug().Where("username=?", fmt.Sprintf("%v", username)).Delete(dao.AccountInfo{})
        //  dao.GDB.Delete(&delAcc)

        c.JSON(http.StatusOK, msg)
    } else {
        msg.Code = 500
        msg.Mess = "can not delete account!"
        c.JSON(http.StatusOK, msg)
    }

}

//更改密码
func ChangepasswdHandler(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获异常:", err)
            msg := ResMsg{
                Code: 500,
                Mess: fmt.Sprintf("%v", err),
            }
            c.JSON(http.StatusOK, msg)
        }
    }()
    msg := ResMsg{
        Code: 200,
        Mess: "return ok",
    }
    username := c.DefaultPostForm("username", "null")
    passwd := c.DefaultPostForm("passwd", "null")
    fmt.Println("更改密码", username, passwd)
    share.Logger.Warn("更改密码账号]..." + fmt.Sprintf("username: %v", username))
    if username != "null" && passwd != "null" {
        // ai := dao.AccountInfo{
        //     Username: username,
        // }

        ai := &dao.AccountInfo{}
        // dao.GDB.Model(&ai).Select("Passwd").Updates(map[string]interface{}{"Passwd": passwd})
        // dao.GDB.Model(&ai{}).Select("Passwd").Updates(map[string]interface{}{"Passwd": passwd})
        dao.GDB.Debug().Model(ai).Where("username = ?", username).Update("passwd", passwd)
        c.JSON(http.StatusOK, msg)
    } else {
        msg.Code = 500
        msg.Mess = "can not hange passwd account!"
        c.JSON(http.StatusOK, msg)
    }
}

////////////////////////////////////////////////////////////添加vpn账号
func VpnUserAddHandler(c *gin.Context) {

    // fmt.Println("liu 是1")

    // vpnliu1 := dao.VpnAccountInfo{}

    // dao.GDB.Debug().Where("username = ?", "liuzhuang").First(&vpnliu1)
    // fmt.Printf("vpnliu1...,%#v", vpnliu1)
    // vpnliu1.VpnAcDoid = "1"
    // dao.GDB.Save(&vpnliu1)

    // fmt.Println("liu 是1、、、、、、、、、、、、")

    // fmt.Println("先进行orm的外键查询......................")
    // vpnliu := dao.VpnAccountInfo{}

    // var DP dao.DomainPermission

    // dao.GDB.Debug().Where("username = ?", "liuzhuang").First(&vpnliu)
    // fmt.Printf("vpnliu...,%#v", vpnliu)
    // dao.GDB.Debug().Model(&vpnliu).Preload("DomainPermissions").Find(&vpnliu)
    // dao.GDB.Debug().Model(&vpnliu).Related(&vpnliu.DomainPermissions).Find(&vpnliu)

    //一对多查询
    // dao.GDB.Debug().Model(vpnliu).Where("username = ?", "liuzhuang").Preload("DomainPermissions").Find(&vpnliu)

    // dao.GDB.Debug().Where("username = ?", "liuzhuang").First(&vpnliu)
    // fmt.Printf("DP...,%#v", DP)
    // fmt.Printf("vpnliu...,%#v", vpnliu)
    // dao.GDB.Debug().Model(&vpnliu).Related(&DP)
    // fmt.Printf("DP:%#v", DP)

    // fmt.Printf("vpnliu...,%#v", vpnliu)
    // for _, item := range vpnliu {
    //     fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>")
    //     fmt.Printf("%v", item.VpnAcDoid)
    //     fmt.Printf("%v", item.DomainPermissions)
    //     fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<")
    // }
    // fmt.Println("                  ")

    // fmt.Printf("%T,%v", vpnliu, vpnliu)
    // for _, item := range vpnliu.DomainPermissions {
    //     fmt.Println(item.Domain)
    // }
    // fmt.Println(len(vpnliu))

    fmt.Println("先进行orm的外键查询...........end..........")

    // dao.GDB.Debug().Model(&dao.VpnAccountInfo).Where("username = ?", "liuzhuang").Update("DomainPermissionID":5)
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获异常:", err)
            errmsg := ResMsg{
                Code: 500,
                Mess: fmt.Sprintf("%v", err),
            }
            c.JSON(http.StatusOK, errmsg)
        }
    }()

    msg := ResMsg{
        Code: 200,
        Mess: "return ok",
    }

    username := c.DefaultPostForm("username", "null")
    passwd := c.DefaultPostForm("passwd", "null")
    fmt.Println("vpn 账号添加：", username, passwd)
    share.Logger.Warn("添加vpn账号:]..." + fmt.Sprintf("username: %v, passwd: %v", username, passwd) + "\n\n")
    if username != "null" && passwd != "null" {

        userinfos := dao.VpnAccountInfo{
            Username: username,
            Passwd:   passwd,
            Stats:    "无登录", //刚创建，没有登录
            // DomainPermissionID: 1,     //默认
            CreateTimes: time.Now().Format("2006-01-02 15:04:05"),
        }

        dao.GDB.Create(&userinfos)

        // 创建vpn用户

        share.Logger.Warn("添加vpn账号成功!...\n\n")
        c.JSON(http.StatusOK, msg)

    } else {

        msg.Code = 500
        msg.Mess = "账号或者密码获取失败！"
        share.Logger.Warn("添加vpn账号失败,账号或者密码获取失败...")
        c.JSON(http.StatusOK, msg)

    }

}

//////////////////////////////////VpnUserlist

func VpnUserlist(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获异常:", err)
            errmsg := ResMsg{
                Code: 500,
                Mess: fmt.Sprintf("%v", err),
            }
            c.JSON(http.StatusOK, errmsg)
        }
    }()

    limits := c.PostForm("limit")
    limitint, _ := strconv.Atoi(limits)
    pages := c.PostForm("page")
    pagesint, _ := strconv.Atoi(pages)

    fmt.Println("limitint,pagesint", limitint, pagesint)

    vpnuserinfos := dao.GetVpnUserinfo()

    // fmt.Printf("%#v", vpnuserinfos)

    share.Logger.Warn("vpn列表展示...")

    // fmt.Println("\n limit,page \n", limits, pages)
    // pagehou := (pagesint * limitint) + 1
    // pagepre := (pagesint * limitint) - limitint + 1

    pagehou := (pagesint * limitint)
    pagepre := limitint * (pagesint - 1)

    // fmt.Println("vpnuserinfos", len(vpnuserinfos))
    // fmt.Println("pagehou...........:", pagehou)
    // fmt.Println("pagepre...........:", pagepre)

    if pagehou > len(vpnuserinfos) {
        pagehou = len(vpnuserinfos)
    }
    c.JSON(http.StatusOK, gin.H{
        "code":  0,
        "msg":   "",
        "count": len(vpnuserinfos),
        // "data":  vpnuserinfos,
        "data": vpnuserinfos[pagepre:pagehou],
    })

}
