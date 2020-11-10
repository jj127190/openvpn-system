package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *sql.DB

var GDB *gorm.DB

//账号数据库

func init() {

	dsn := "root:8927126@tcp(127.0.0.1:3306)/openvpensystem"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("初始化数据库失败......, err:%v\n", err)
		return
	}

	//gorm
	gdb, err := gorm.Open("mysql", "root:8927126@(172.30.0.196:22809)/openvpensystem?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("gorm conn is fail...")
		panic(err)
	}
	gdb.SingularTable(true)
	gdb.AutoMigrate(&AccountInfo{}, &VpnAccountInfo{}, &DomainPermission{}, &PermissionDisgroup{})

	GDB = gdb
	//gorm
}

func initDB(dsn string) (err error) {
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open数据库失败......, err:%v\n", err)
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(50)
	return nil
}

func QueryPass(username string) (string, error) {
	// 查询单行数据
	var user User

	sqlStr := fmt.Sprintf("select passwd from account_info where username=\"%s\"", username)
	err := DB.QueryRow(sqlStr).Scan(&user.Password)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return "", err
	}

	fmt.Printf("查询结果：%#v\n", user)
	return user.Password, nil
}
