package main

import (
	"opsystem/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Router("/user/add", &controllers.CommonController{}, "post:AddUser")
	beego.Router("/user/delete", &controllers.UserController{}, "post:DeleteUser")
	beego.Router("/user/update", &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/user/getjurieses", &controllers.UserController{}, "get:GetJuris")
	beego.Router("/user/gettotalinfo", &controllers.UserController{}, "get:GetTotalInfo")
	beego.Router("/login", &controllers.CommonController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "post:LoginOut")
	beego.Router("/checkname", &controllers.CommonController{}, "post:CheckUsername")
	beego.Router("/user/getuserinfobyid", &controllers.CommonController{}, "get:GetUserInfoById")

	beego.Router("/product/add", &controllers.ProductController{}, "post:AddProduct")
	beego.Router("/product/update", &controllers.ProductController{}, "put:Update")
	beego.Router("/product/all", &controllers.ProductController{}, "get:GetAll")
	beego.Router("/product/delete", &controllers.ProductController{}, "delete:Delete")
	beego.Router("/product/get", &controllers.ProductController{}, "post:GetProdsByPage")
	beego.Router("/product/count", &controllers.ProductController{}, "get:GetCount")
	beego.Run()
}
