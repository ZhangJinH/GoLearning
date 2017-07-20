package main

import (
	"qq/controllers"
	_ "qq/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/qq")
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Router("/v1/user/login", &controllers.UserController{}, "post:Login")
	beego.Run()
}
