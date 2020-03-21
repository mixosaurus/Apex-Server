package main

import (
	"com.vic/Controller"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

type UserInfo struct {
	UserName string
	Address  string
	Sex      byte
	Company  Company
}

type Company struct {
	Name       string
	EmployeeId int
	Department []string
}

func main() {
	app := iris.New()
	controller.Register(app)                                                   //方法注册
	app.Run(iris.Addr(":8888"), iris.WithoutServerError(iris.ErrServerClosed)) //启动服务器
	//Test()
}

func Test() {

	config := viper.New()
	//config.SetConfigFile("config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")

	////windows环境下为%GOPATH，linux环境下为$GOPATH
	//config.AddConfigPath("/Users/yangyue/workspace/go/src/webDemo/")
	////设置配置文件类型
	//config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	//fmt.Printf("userName:%s sex:%s company.name:%s \n", config.Get("userName"), config.Get("sex"), config.Get("company.name"))

	//也可以直接反序列化为Struct

	var userInfo UserInfo
	if err := config.Unmarshal(&userInfo); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println("userInfo: ", userInfo)

}
