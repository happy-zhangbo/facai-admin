package services

import (
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/repositories"
)

type OrderService interface {
	GetOrderList(m map[string]interface{}) (result datamodel.Result)
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

var repo = repositories.NewOrderRepo()

func (o orderService) GetOrderList(m map[string]interface{}) (result datamodel.Result) {
	var repo = repositories.NewOrderRepo()
	total, orderList := repo.GetOrderList(m)
	maps := make(map[string]interface{}, 2)
	maps["Total"] = total
	maps["Order"] = orderList
	result.Data = maps
	result.Code = 0
	result.Msg = "SUCCESS"

	return
}
