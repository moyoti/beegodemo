package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) URLMapping() {
	c.Mapping("Test", c.Test)
}

// @router /tt [get]
func (c *TestController) Test() {
	info := c.Input().Get("testInfo")
	c.Data["json"] = info
	c.ServeJSON()
}
