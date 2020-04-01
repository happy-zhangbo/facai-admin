package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/yaboyou/facai-admin/web/controllers"
	"github.com/yaboyou/facai-admin/web/middleware"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl)).Handle(controllers.NewLoginController())
	app.Use(middleware.GetJWT().Serve) // jwt
	mvc.New(app.Party(bathUrl + "/user")).Handle(controllers.NewUserInfoController())
	mvc.New(app.Party(bathUrl + "/order")).Handle(controllers.NewOrderController())

}
