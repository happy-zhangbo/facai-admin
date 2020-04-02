package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/services"
	"log"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{Service: services.NewProductService()}
}

func (p *ProductController) PostList() (result datamodel.Result) {
	var m map[string]interface{}
	err := p.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["page"] == "" || m["page"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 page"
		return
	}
	if cast.ToUint(m["page"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 page"
		return
	}
	if m["size"] == "" || m["size"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 size"
		return
	}
	if cast.ToUint(m["size"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 size"
		return
	}
	return p.Service.GetProductList(m)
}

func (p *ProductController) PostSave() (result datamodel.Result) {
	var product datamodel.Product
	if err := p.Ctx.ReadJSON(&product); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}
	fmt.Println(product)
	return p.Service.SaveProduct(product)
}
func (p *ProductController) PostGet() (result datamodel.Result) {
	var m map[string]interface{}
	err := p.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["id"] == "" || m["id"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 id"
		return
	}
	if cast.ToUint(m["id"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 id"
		return
	}
	return p.Service.GetProduct(cast.ToUint(m["id"]))
}
func (p *ProductController) PostDel() (result datamodel.Result) {
	var m map[string]interface{}
	err := p.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["id"] == "" || m["id"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 id"
		return
	}
	if cast.ToUint(m["id"]) == 0 {
		result.Code = -1
		result.Msg = "参数错误 id"
		return
	}
	return p.Service.DeleteProduct(cast.ToUint(m["id"]))
}
