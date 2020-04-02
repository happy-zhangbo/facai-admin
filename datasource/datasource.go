package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yaboyou/facai-admin/config"
	"github.com/yaboyou/facai-admin/datamodel"
	"strings"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	path := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	//Createtable();
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {

		str := strings.Split(defaultTableName, "_")
		fmt.Println(str)
		if len(str) <= 1 {
			return defaultTableName
		}
		fmt.Println(str[0] + Capitalize(str[1]))
		return str[0] + Capitalize(str[1])
	}
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 初始化表 如果不存在该表 则自动创
func Createtable() {
	GetDB().AutoMigrate(
		&datamodel.Address{},
		&datamodel.Cart{},
		&datamodel.Order{},
		&datamodel.OrderDetail{},
		&datamodel.Product{},
		&datamodel.ProductType{},
		&datamodel.Userinfo{},
	)
}
