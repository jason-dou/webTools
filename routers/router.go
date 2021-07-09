package routers

import (
	"github.com/beego/beego/v2/server/web"
	"webTools/controllers"
)

func init() {
	web.Router("/", &controllers.IndexController{})
    web.Router("/calendar", &controllers.CalendarController{})
}
