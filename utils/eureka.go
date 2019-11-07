package utils

import (
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"fmt"
	"time"
)

var client *eureka.Client

func GetEureka() *eureka.Client {
	if client == nil {
		initEureka()
	}
	return client
}
func RegisterInstance() {

	instance := eureka.NewInstanceInfo("10.106.159.222", "GO-TEST", "10.106.159.222", 8080, 30, false) //Create a new instance to register
	//instance.Metadata = &eureka.MetaData{
	//	Map: make(map[string]string),
	//}
	//instance.Metadata.Map["HomePageUrl"] = "bar"            //add metadata for example
	//instance.Metadata.Map["StatusPageUrl"] = "bar"          //add metadata for example
	err := client.RegisterInstance("GO-TEST", instance) // Register new instance in your eureka(s)
	if err != nil {
		fmt.Printf("RegisterInstance err %s\n", err.Error())
		return
	}
	client.GetApplication(instance.App)                   // retrieve the application "test"
	client.GetInstance(instance.App, instance.HostName)   // retrieve the instance from "test.com" inside "test"" app
	// say to eureka that your app is alive (here you must send heartbeat before 30 sec);
	if err:=client.SendHeartbeat(instance.App, instance.HostName); err!=nil{
		fmt.Printf("SendHeartbeat err %s\n", err.Error())
		return
	}

	t:=time.NewTicker(time.Second*10)
	for {
		select {
			case <-t.C :
				client.SendHeartbeat(instance.App, instance.HostName)
		}
	}
}
func initEureka() {
	client = eureka.NewClient([]string{
		"http://fortune-app-eureka.154.test.yirendai.com/eureka", //From a spring boot based eureka server
		// add others servers here
	})
	client.Config.DialTimeout = time.Second * 10

	//apps, err := client.GetApplications()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, appinfo := range apps.Applications {
	//	fmt.Println(appinfo.Name)
	//	for _, ins := range appinfo.Instances {
	//		fmt.Printf("%s:%+v\n", appinfo.Name, ins)
	//	}
	//}

	//appid := "GATEWAY"
	//client.GetApplication(appid)
}

