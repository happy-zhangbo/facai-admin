package repositories

import (
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/datasource"
)

type OrderReop interface {
	GetOrderList(m map[string]interface{}) (totle int, orderList []datamodel.Order)
	GetOrderDetail(m map[string]interface{}) (orderDetail []datamodel.OrderDetail)
}

func NewOrderRepo() OrderReop {
	return &orderRepo{}
}

type orderRepo struct{}

func (o orderRepo) GetOrderDetail(m map[string]interface{}) (orderDetail []datamodel.OrderDetail) {
	var db = datasource.GetDB()
	err := db.Where(&datamodel.OrderDetail{OdOid: cast.ToInt(m["oid"])}).Find(&orderDetail).Error
	if err != nil {
		panic("select Error")
	}
	for i := 0; i < len(orderDetail); i++ {
		db.Model(orderDetail[i]).Related(&orderDetail[i].ProductSpecs, "OdPsid")
		ps := orderDetail[i].ProductSpecs
		for j := 0; j < len(ps); j++ {
			db.Model(ps[j]).Related(&ps[j].Product, "SProductid")
		}
	}
	return
}

func (o orderRepo) GetOrderList(m map[string]interface{}) (totle int, orderList []datamodel.Order) {
	var db = datasource.GetDB()
	db.Table("order").Count(&totle)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Order("o_createtime desc").Find(&orderList).Error
	if err != nil {
		panic("select Error")
	}
	//for i := 0;i < len(orderList) ;i++  {
	//	db.Model(&orderList[i]).Related(&orderList[i].OrderDetail,"OdOid")
	//}
	return
}
