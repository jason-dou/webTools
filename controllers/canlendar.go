package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type CalendarController struct {
	web.Controller
}

func (c *CalendarController) Get() {
	c.Data["Website"] = "website.me"
	c.Data["Email"] = "email@gmail.com"
	c.TplName = "calendar.tpl"
}
