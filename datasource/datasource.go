package datasource

import (
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
