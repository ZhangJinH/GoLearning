package controllers

import (
	"encoding/json"
	"fmt"
	"opsystem/model"
	"strconv"

	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) AddUser() {
	var v model.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Jid = 1
		if err := model.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

type LoginInfo struct {
	Username string
	Password string
}

func (c *CommonController) Login() {
	var v LoginInfo
	var res map[string]string
	res = make(map[string]string)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := model.GetUserByUsername(v.Username)
		if err != nil {
			res["code"] = "1000"
			res["msg"] = "账号不存在"
			c.Data["json"] = res
		} else if user.Password != v.Password {
			res["code"] = "1001"
			res["msg"] = "密码错误"
			c.Data["json"] = res
		} else {
			fmt.Printf("the user is: %v", user)
			res["code"] = "1002"
			res["msg"] = "登陆成功"
			idStr := strconv.Itoa(user.Id)
			res["id"] = idStr
			c.Data["json"] = res
			c.SetSession("uid", user.Id)
			c.SetSession("jid", user.Jid)
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *CommonController) CheckUsername() {
	username := c.GetString("username")
	_, err := model.GetUserByUsername(username)
	var res map[string]string
	res = make(map[string]string)
	if err != nil {
		c.Data["json"] = nil
	} else {
		res["code"] = "1003"
		res["msg"] = "账号已存在"
		c.Data["json"] = res
	}
	c.ServeJSON()
}

func (c *CommonController) GetUserInfoById() {
	id := c.GetSession("uid").(int)
	v, err := model.GetUserInfo(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *CommonController) Upload() {
	f, file, err := c.GetFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}
	var attachment string
	attachment = file.Filename
	dirPath := "./upload/" + attachment
	f.Close()
	err = c.SaveToFile("file", dirPath)
	if err != nil {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
