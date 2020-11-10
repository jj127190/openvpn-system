package registry

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
)


var ConsulClient *consul.Client

func init() {
	config := consul.DefaultConfig()
	config.Address = "172.30.0.196:8500"

	client, err := consul.NewClient(config)

	if err != nil {
		fmt.Println("err:", err)
	}
	ConsulClient = client
}

//注册
func RegService() {
	//config := consul.DefaultConfig()
	//config.Address="172.30.0.196:8500"

	reg := consul.AgentServiceRegistration{}
	reg.ID = "newServerID"
	reg.Name = "newService"
	reg.Address = "172.30.0.196"
	reg.Port = 8088
	reg.Tags = []string{"newTest"}

	check := consul.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://172.30.0.196:8088/health"
	check.DeregisterCriticalServiceAfter = "30s"            //30s失败后自动将注册服务删除
	reg.Check = &check
	//client,err := consul.NewClient(config)

	//if err != nil{
	//	fmt.Println("err:",err)
	//}

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		fmt.Println("err:", err)
	}
}

//反注册
func UnRegservice() {
	ConsulClient.Agent().ServiceDeregister("userServicelmx") //传入id
}








//服务发现


// package main
 
// import (
// 	"fmt"
// 	consulapi "github.com/hashicorp/consul/api"
// )
 
// const (
// 	consulAgentAddress = "172.30.0.196:8500"
// )
 
// // 从consul中发现服务
// func ConsulFindServer()  {
// 	// 创建连接consul服务配置
// 	config := consulapi.DefaultConfig()
// 	config.Address = consulAgentAddress
// 	client, err := consulapi.NewClient(config)
// 	if err != nil {
// 		fmt.Println("consul client error : ", err)
// 	}
 
// 	// 获取指定service
// 	service, _, err := client.Agent().Service("userService", nil)
// 	if err == nil{
// 		fmt.Println("服务发现....")
// 		fmt.Println(service.Address)
// 		fmt.Println(service.Port)
// 	}
 
//     	//只获取健康的service
// 	//serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
// 	//if err == nil{
// 	//	fmt.Println(serviceHealthy[0].Service.Address)
// 	//}
 
// }
 
// func main()  {
// 	ConsulFindServer()
// }