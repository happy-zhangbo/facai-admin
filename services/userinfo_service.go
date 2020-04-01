package services

import (
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/repositories"
	"github.com/yaboyou/facai-admin/utils"
	"github.com/yaboyou/facai-admin/web/middleware"
)

type UserInfoService interface {
	Login(m map[string]string) (result datamodel.Result)
	GetUserList(m map[string]interface{}) (result datamodel.Result)
}

type userInfoService struct {
}

func NewUserInfoService() UserInfoService {
	return &userInfoService{}
}

func (u userInfoService) Login(m map[string]string) (result datamodel.Result) {

	if m["username"] == "" {
		result.Code = -1
		result.Msg = "请输入用户名！"
		return
	}
	if m["password"] == "" {
		result.Code = -1
		result.Msg = "请输入密码！"
		return
	}
	user := datamodel.Userinfo{}
	if "zhangbo" == m["username"] && utils.GetMD5String("zhangbo") == utils.GetMD5String(m["password"]) {
		user.Token = middleware.GenerateToken(user)
		user.UNickname = "张博"

		result.Code = 0
		result.Data = user
		result.Msg = "登录成功"
		return
	} else {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
		return
	}
}

func (u userInfoService) GetUserList(m map[string]interface{}) (result datamodel.Result) {
	var repo = repositories.NewUserInfoRepository()
	total, userList := repo.GetUserList(m)
	maps := make(map[string]interface{}, 2)
	maps["Total"] = total
	maps["Users"] = userList
	result.Data = maps
	result.Code = 0
	result.Msg = "SUCCESS"
	return
}
