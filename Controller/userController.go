package controller

import (
	"com.vic/Handler"
	"github.com/kataras/iris"
)

//func welcome

func GetUserInfo(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	user := Handler.GetUserInfo(userId)
	ctx.JSON(user)
}

//func GetPosts(ctx iris.Context){
//	userId:=ctx.URLParam("userId")
//	rows:=Handler.GetPosts(userId)
//	ctx.JSON(rows)
//}
