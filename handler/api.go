package handler

import (
	"VpnAudit/dao"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"reflect"
	"strconv"
)

func StrctToSlice(f dao.AccountInfo) []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}

func AdminlistHandler(c *gin.Context) {

	fmt.Println("AdminlistHandler...................")
	// limits := c.PostForm("limit")
	// limitint, _ := strconv.Atoi(limits)
	// pages := c.PostForm("page")
	// pagesint, _ := strconv.Atoi(pages)

	//c.PostForm("limit")
	//strconv.Atoi(limits)

	userinfos := dao.GetAccountinfo()

	//fmt.Printf("%#v", userinfos)

	// resdataslice := make([]map[string]interface{}, 10)

	// for _, item := range userinfos {
	// 	resmap := make(map[string]interface{}, 10)
	// 	j, _ := json.Marshal(item)
	// 	json.Unmarshal(j, &resmap)
	// 	resdataslice = append(resdataslice, resmap)
	// }
	// fmt.Println("0000")
	// fmt.Printf("%#v", resdataslice)

	// fmt.Println("pagesint", pagesint)
	// pagehou := (pagesint * limitint) + 1
	// pagepre := (pagesint * limitint) - limitint + 1

	// fmt.Println("pagepre:", pagepre)
	// fmt.Println("pagehou", pagehou)
	// fmt.Printf("%T", pagepre)

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": len(userinfos),
		"data":  userinfos,
		// "data": userinfos[pagepre:pagehou],
	})
}

func TestHandler(c *gin.Context) {
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//fmt.Println("-----------")
	//fmt.Println(string(buf[0:n]))
	limits := c.PostForm("limit")
	limitint, _ := strconv.Atoi(limits)
	pages := c.PostForm("page")
	pagesint, _ := strconv.Atoi(pages)
	fmt.Println("=============")
	fmt.Printf("%T", limitint)
	var dataslice []map[string]string
	// var datasline []map[string]interface{}
	c.PostForm("limit")
	strconv.Atoi(limits)

	// limitint := strconv.Itoa(limits)
	for i := 0; i < 100; i++ {
		for i := 0; i < 100; i++ {
			datas := make(map[string]string, 10)
			datas["id"] = fmt.Sprintf("%d", i)
			datas["email"] = "xianxin@layui.com"

		}
		//   [{k,k,k,v}]
		datas := make(map[string]string, 10)
		datas["username"] = "saber"
		datas["id"] = fmt.Sprintf("%d", i)
		datas["email"] = "xianxin@layui.com"
		datas["sex"] = "男"
		datas["city"] = "浙江"
		datas["sign"] = "sign conent 内容"
		datas["experience"] = "116"
		datas["ip"] = "192.168.0.8"
		datas["logins"] = "108"
		datas["joinTime"] = "2016-10-14"

		dataslice = append(dataslice, datas)
	}
	fmt.Println("pagesint", pagesint)
	pagehou := (pagesint * limitint) + 1
	pagepre := (pagesint * limitint) - limitint + 1
	fmt.Println("pagepre:", pagepre)
	fmt.Println("pagehou", pagehou)
	fmt.Printf("%T", pagepre)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": len(dataslice),
		"data":  dataslice[pagepre:pagehou],
	})
}
