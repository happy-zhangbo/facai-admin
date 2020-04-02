package repositories

import (
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/datasource"
)

type UserinfoRepo interface {
	GetUserList(m map[string]interface{}) (total int, userinfo []datamodel.Userinfo)
}

func NewUserInfoRepository() UserinfoRepo {
	return &userInfoRepo{}
}

type userInfoRepo struct{}

var db = datasource.GetDB()

func (u userInfoRepo) GetUserList(m map[string]interface{}) (total int, userinfo []datamodel.Userinfo) {

	db.Table("userinfo").Count(&total)

	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&userinfo).Error
	//err := db.Select("u_id").Find(&userinfo).Error
	if err != nil {
		panic("select Error")
	}
	return
}
