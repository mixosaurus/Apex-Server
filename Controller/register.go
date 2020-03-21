package controller

import "github.com/kataras/iris"

func Register(app *iris.Application) {
	userController := app.Party("/user")

	{
		userController.Get("/getUserInfo/", GetUserInfo)
	}

}
