package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	// cid, _ := c.GetSession("uid").(int)
	// jid, _ := c.GetSession("jid").(int)
	// fmt.Printf("the login user's uid is:%d\n", cid)
	// if cid < 1 || jid < 1 {
	// 	c.CustomAbort(http.StatusUnauthorized, "{\"state\":\"prerror\"}")
	// 	return
	// }
}
