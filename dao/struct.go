package dao

import (
	"github.com/jinzhu/gorm"
)

// vpn 平台账号
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

//账号 -> 组 -> 域名

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Createtime string `json:"createtime"`
	Lastloggin string `json:"lastloggin"`
	Lastlogout string `json:"lastlogout"`
	Nowstatus  string `json:"nowstatus"`
	Level      string `json:"level"`
}
