package services

import (
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/repositories"
)

type ProductService interface {
	GetProductList(m map[string]interface{}) (result datamodel.Result)
	SaveProduct(product datamodel.Product) (result datamodel.Result)
	DeleteProduct(id uint) (result datamodel.Result)
	GetProduct(id uint) (result datamodel.Result)
}
type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

var repoProduct = repositories.NewProductRepo()

func (p productService) GetProductList(m map[string]interface{}) (result datamodel.Result) {
	total, productList := repoProduct.GetProductList(m)
	maps := make(map[string]interface{}, 2)
	maps["Total"] = total
	maps["Product"] = productList
	result.Data = maps
	result.Code = 0
	result.Msg = "SUCCESS"
	return
}

func (p productService) SaveProduct(product datamodel.Product) (result datamodel.Result) {
	err := repoProduct.SaveProduct(product)
	if err != nil {
		result.Code = -1
		result.Msg = "保存失败"
	} else {
		result.Code = 1
		result.Msg = "保存成功"
	}
	return
}

func (p productService) DeleteProduct(id uint) (result datamodel.Result) {

	err := repoProduct.DeleteProduct(id)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}

func (p productService) GetProduct(id uint) (result datamodel.Result) {
	product, err := repoProduct.GetProduct(id)
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
	} else {
		result.Data = product
		result.Code = 0
		result.Msg = "SUCCESS"
	}
	return
}
