package controllers

import (
	"encoding/json"
	"opsystem/model"
	"strconv"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) AddProduct() {
	var product model.Product
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &product); err == nil {
		if err := model.AddProduct(&product); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *ProductController) GetAll() {
	prods, err := model.GetAllProds()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = prods
	}
	c.ServeJSON()
}

func (c *ProductController) GetCount() {
	total, err := model.GetProdsCount()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = total
	}
	c.ServeJSON()
}

func (c *ProductController) GetProdsByPage() {
	currentStr := c.GetString("current")
	pagesizeStr := c.GetString("pagesize")
	current, _ := strconv.Atoi(currentStr)
	pagesize, _ := strconv.Atoi(pagesizeStr)
	if prods, err := model.GetProdsByPage(current, pagesize); err == nil {
		c.Data["json"] = prods
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *ProductController) Update() {
	var v *model.Product
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := model.UpdateProd(v); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *ProductController) Delete() {
	idStr := c.GetString("pid")
	pid, _ := strconv.Atoi(idStr)
	if err := model.DeleteProd(pid); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
