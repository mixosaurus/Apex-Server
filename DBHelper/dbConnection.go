package DBHelper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DBConnectionParm struct {
	UserName string
	DBServer string
	PWD      string
	DBName   string
}

func NewInstance() *gorm.DB {

	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	//config.SetDefault("ContentDir", "content")
	config.ReadInConfig() //读取配置文件至config
	var configData DBConnectionParm
	config.Unmarshal(&configData)

	//读取配置文件至configData
	if err := config.Unmarshal(&configData); err != nil {
		fmt.Println(err.Error())
	}

	//decryptedPwd := AesHelper.AESDecrypt(configData.PWD)

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%v:%v@(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			configData.UserName,
			configData.PWD,
			configData.DBServer,
			configData.DBName,
		))

	//db, err := gorm.Open("mysql","root:abcdefg@(127.0.0.1:3306)/apex_db?charset=utf8mb4&parseTime=True&loc=Local")

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	if err != nil {
		fmt.Println(err)
	}
	return db
}
