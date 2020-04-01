package repositories

import (
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/datasource"
)

type OrderReop interface {
	GetOrderList(m map[string]interface{}) (totle int, orderList []datamodel.Order)
}

func NewOrderRepo() OrderReop {
	return &orderRepo{}
}

type orderRepo struct{}

func (o orderRepo) GetOrderList(m map[string]interface{}) (totle int, orderList []datamodel.Order) {
	var db = datasource.GetDB()
	db.Table("order").Count(&totle)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&orderList).Error
	if err != nil {
		panic("select Error")
	}
	return
}
