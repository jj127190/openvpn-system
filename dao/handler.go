package dao

import (
	// "fmt"
	"strings"
)

func GetAccountinfo() []AccountInfo {
	var ais []AccountInfo
	GDB.Find(&ais)
	return ais
}

//vpn账号
func GetVpnUserinfo() []VpnAccountInfo {
	var va []VpnAccountInfo
	var pd []PermissionDisgroup
	pdmaps := make(map[uint]string, 10)
	GDB.Find(&va)
	GDB.Find(&pd)

	for _, item := range pd {
		pdmaps[item.VpnAcDoid] = item.GroupName
	}

	for index := range va {

		// fmt.Println(va[index], va[index].Username)
		va[index].Passwd = pdmaps[va[index].VpnAcDoid]
	}
	// fmt.Println(pdmaps)
	// VpnAcDoid          uint
	// 	PermissionDisgroup

	// var vai []VpnAccountInfo
	// GDB.Find(&vai)
	// GDB.Model(&vai).Preload("PermissionDisgroup").Find(&vai)
	// fmt.Println("vai..........vai.........", vai[0])
	// fmt.Println("...........................", va)
	return va
}

func GetDomain() []DomainPermission {
	var dm []DomainPermission
	GDB.Find(&dm)
	return dm
}

//获取权限组
func GetDomainGroup() []PermissionDisgroup {
	var pd []PermissionDisgroup
	GDB.Find(&pd)
	pdres := GetDomainConn(pd)
	return pdres
}

//根据组获取域名
func GetDomainConn(pd []PermissionDisgroup) []PermissionDisgroup {
	for index, item := range pd {
		// fmt.Println("关联域名......")
		// fmt.Printf("%T,%v", pd, pd)
		// fmt.Println("Perdomk......")
		// fmt.Printf("%T,%v", item.Perdomk, item.Perdomk)
		res := strings.Split(item.Perdomk, ",")
		// fmt.Println("res.....", res)
		// fmt.Printf("%T,%v", res)
		var cdomains string
		for _, items := range res {
			var dp DomainPermission
			GDB.Where("Perdomk = ?", items).Find(&dp)
			// fmt.Println("dp:....", dp.Domain)
			if len(cdomains) != 0 {
				cdomains = cdomains + ","
			}
			cdomains = cdomains + dp.Domain
		}
		// fmt.Println("cdomains", cdomains)
		// fmt.Printf("%T,%v", pd[index], pd[index])
		pd[index].Perdomk = cdomains
	}
	return pd

}

func PerGroupSave(Groupname, SelectCon string) {

	var PD PermissionDisgroup
	PD = PermissionDisgroup{GroupName: Groupname, Perdomk: SelectCon}
	GDB.Create(&PD)
	PD.VpnAcDoid = PD.ID
	GDB.Save(&PD)

}

//编辑权限组
func PerGroupEdit(Groupname, SelectCon string) {

	GDB.Debug().Model(&PermissionDisgroup{}).Where("group_name = ?", Groupname).Update("perdomk", SelectCon)

}

// fun waijian(){
// 	vpnliu := VpnAccountInfo{
//         Username: "liuzhuang",
//     }

//     dao.GDB.Model(&vpnliu).Updates(VpnAccountInfo{DomainPermissionID:1})

//     DP := dao.DomainPermission{}

//     dao.GDB.Debug().Model(&vpnliu).Related(&DP)
//     fmt.Printf("DP:%#v", DP)
// }

func VpnAccBindGroup(VpnAccount, DomainGroup string) {

	var vis VpnAccountInfo
	var pd PermissionDisgroup
	var viss []VpnAccountInfo

	// GDB.Find(&vis)
	GDB.Debug().Where("username = ?", VpnAccount).First(&vis)

	GDB.Debug().Where("group_name = ?", DomainGroup).First(&pd)

	viss = append(viss, vis)
	pd.VpnAccountInfo = viss
	GDB.Save(&pd)

	// var vpnliu []VpnAccountInfo
	//外键查询
	// GDB.Where("username = ?", "liuzhuang").Find(&vpnliu)
	// GDB.Model(&vpnliu).Preload("PermissionDisgroup").Find(&vpnliu)
	// fmt.Println("vpnliu..........vpndp.........", vpnliu)

}

func PerGroupDel(id string) {

	GDB.Debug().Where("id = ?", id).Delete(&PermissionDisgroup{})
}
