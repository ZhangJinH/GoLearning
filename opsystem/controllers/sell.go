package controllers

import (
	"encoding/json"
	"opsystem/model"
	"strconv"
)

type SellController struct {
	BaseController
}

func (c *SellController) AddSell() {
	var v model.Sell
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := model.AddSell(&v); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SellController) CheckAlreadySet() {
	var v model.Sell
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		has, err := model.CheckAlreadySet(&v)
		if err != nil {
			c.Data["json"] = err.Error()
		} else if has {
			c.Data["json"] = "当前已设置当天当地销量"
		} else {
			c.Data["json"] = nil
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SellController) GetCount() {
	total, err := model.GetSellCount()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = total
	}
	c.ServeJSON()
}

func (c *SellController) GetSellByPage() {
	currentStr := c.GetString("current")
	pagesizeStr := c.GetString("pagesize")
	current, _ := strconv.Atoi(currentStr)
	pagesize, _ := strconv.Atoi(pagesizeStr)
	if sells, err := model.GetSellByPage(current, pagesize); err == nil {
		c.Data["json"] = sells
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SellController) Delete() {
	idStr := c.GetString("id")
	id, _ := strconv.Atoi(idStr)
	err := model.DeleteSell(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "ok"
	}
	c.ServeJSON()
}

func (c *SellController) Update() {
	var v model.Sell
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := model.UpdateSell(&v); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SellController) GetTotalSellNums() {
	if nums, err := model.GetTotalSellNums(); err == nil {
		c.Data["json"] = nums
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *SellController) GetTotalSellPlace() {
	idStr := c.GetString("pid")
	id, _ := strconv.Atoi(idStr)
	if nums, err := model.GetTotalSellPlace(id); err == nil {
		c.Data["json"] = nums
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
