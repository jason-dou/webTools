package routers

import (
	"github.com/beego/beego/v2/server/web"
	"webTools/controllers"
)

func init() {
	web.Router("/", &controllers.IndexController{})
    web.Router("/calendar", &controllers.CalendarController{}, "get:GetAll;post:CreateCalendarEvent")
	web.Router("/calendar/:id", &controllers.CalendarController{}, "get:GetById;delete:DeleteCalendarEvent")
}
