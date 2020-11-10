package handler

import (
	"openvpn-system/dao"

	"fmt"
	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	// "reflect"
	"openvpn-system/common/share"
	"strconv"
)

func DomainShow(c *gin.Context) {

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

	// fmt.Println("limitint,pagesint", limitint, pagesint)

	domainGroups := dao.GetDomainGroup()

	share.Logger.Warn("域名组创建列表展示...")

	// fmt.Println("\n limit,page \n", limits, pages)
	// pagehou := (pagesint * limitint) + 1
	// pagepre := (pagesint * limitint) - limitint + 1

	pagehou := (pagesint * limitint)
	pagepre := limitint * (pagesint - 1)

	// fmt.Println("vpnuserinfos", len(vpnuserinfos))
	// fmt.Println("pagehou...........:", pagehou)
	// fmt.Println("pagepre...........:", pagepre)

	if pagehou > len(domainGroups) {
		pagehou = len(domainGroups)
	}

	fmt.Println(domainGroups)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": len(domainGroups),
		// "data":  vpnuserinfos,
		"data": domainGroups[pagepre:pagehou],
	})

}

func AddPerGroup(c *gin.Context) {
	groupname := c.DefaultPostForm("groupname", "null")
	selectCon := c.DefaultPostForm("selectCon", "null")
	fmt.Println("后台添加域名组..............................")
	fmt.Println(groupname, selectCon)
	if groupname != "null" && selectCon != "null" {
		fmt.Println(groupname, selectCon)
		dao.PerGroupSave(groupname, selectCon)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": "参数传递有问题,请再次检查!",
		})
	}
}

func GetDomains(c *gin.Context) {
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

	// limits := c.PostForm("limit")
	// limitint, _ := strconv.Atoi(limits)
	// pages := c.PostForm("page")
	// pagesint, _ := strconv.Atoi(pages)

	// fmt.Println("limitint,pagesint", limitint, pagesint)

	domains := dao.GetDomain()

	// fmt.Printf("%#v", vpnuserinfos)

	// fmt.Println("\n limit,page \n", limits, pages)
	// pagehou := (pagesint * limitint) + 1
	// pagepre := (pagesint * limitint) - limitint + 1
	// pagehou := (pagesint * limitint)
	// pagepre := limitint * (pagesint - 1)
	// fmt.Println("vpnuserinfos", len(vpnuserinfos))
	// fmt.Println("pagehou...........:", pagehou)
	// fmt.Println("pagepre...........:", pagepre)
	// if pagehou > len(vpnuserinfos) {
	// 	pagehou = len(vpnuserinfos)
	// }
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": domains,
	})

}

func Winitgedomains(c *gin.Context) {
	PD := dao.GetDomainGroup()
	var resms []map[string]string
	for _, item := range PD {
		imap := make(map[string]string, 10)
		imap[item.GroupName] = item.Perdomk
		resms = append(resms, imap)
	}
	// fmt.Println(PD)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": PD,
	})
}

func AddGroupajax(c *gin.Context) {
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
	fmt.Println("ajax 添加组")

	VpnAccount := c.DefaultPostForm("VpnAccount", "null")
	DomainGroup := c.DefaultPostForm("DomainGroup", "null")
	if DomainGroup != "null" && VpnAccount != "null" {
		fmt.Println(VpnAccount, DomainGroup)

		dao.VpnAccBindGroup(VpnAccount, DomainGroup)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})

	}

}

//权限组删除
func Pergroupdel(c *gin.Context) {
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
	fmt.Println("权限组删除")
	groupid := c.DefaultPostForm("gid", "null")

	if groupid != "null" {
		fmt.Printf("%T,%v", "权限组id........", groupid, groupid)
		dao.PerGroupDel(groupid)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "删除成功",
	})

}

func EditPerGroup(c *gin.Context) {
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
	fmt.Println("权限组编辑")

	groupname := c.DefaultPostForm("groupname", "null")
	selectCon := c.DefaultPostForm("selectCon", "null")
	fmt.Println("后台编辑域名组..............................")
	fmt.Println(groupname, selectCon)
	if groupname != "null" && selectCon != "null" {

		dao.PerGroupEdit(groupname, selectCon)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": "参数传递有问题,请再次检查!",
		})
	}

}
