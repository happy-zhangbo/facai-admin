package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/services"
	"github.com/yaboyou/facai-admin/utils"
	"log"
)

type UserInfoController struct {
	Ctx     iris.Context
	Service services.UserInfoService
}

func NewUserInfoController() *UserInfoController {
	return &UserInfoController{Service: services.NewUserInfoService()}
}

func (u *UserInfoController) PostList() (result datamodel.Result) {
	var m map[string]interface{}
	err := u.Ctx.ReadJSON(&m)
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
	return u.Service.GetUserList(m)
}

func (u *UserInfoController) GetToken() (result datamodel.Result) {
	maps := make(map[string]interface{}, 2)
	maps["Token"] = utils.CreateUptoken()
	maps["Url"] = "http://image.lonelysky.com.cn"
	result.Data = maps
	result.Code = 0
	result.Msg = "SUCCESS"
	return result
}
