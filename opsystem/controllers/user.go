package controllers

import (
	"encoding/json"
	"opsystem/model"

	"strconv"
)

type UserController struct {
	BaseController
}

func (c *UserController) DeleteUser() {
	idStr := c.GetString("id")
	id, _ := strconv.Atoi(idStr)
	err := model.DeleteUSer(id)
	if err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UserController) UpdateUser() {
	var v model.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := model.UpdateUser(&v); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UserController) GetJuris() {
	jurieses, err := model.GetJuris()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = jurieses
	}
	c.ServeJSON()
}

func (c *UserController) GetTotalInfo() {
	users, err := model.GetTotalUserInfo()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

func (c *UserController) LoginOut() {
	c.DestroySession()
	c.Ctx.Output.SetStatus(401)
}
