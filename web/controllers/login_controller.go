package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/services"
	"log"
)

type LoginController struct {
	Ctx     iris.Context
	Service services.UserInfoService
}

func NewLoginController() *LoginController {
	return &LoginController{Service: services.NewUserInfoService()}
}

func (g *LoginController) PostLogin() datamodel.Result {
	var m map[string]string
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}

	result := g.Service.Login(m)
	return result
}
