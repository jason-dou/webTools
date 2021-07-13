package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"webTools/managers"
	"webTools/models"
)

type CalendarEventResponse struct {
	CalendarEvents []models.CalendarEvent `json:"calendar_events"`
}

type CalendarController struct {
	web.Controller
}

func (cc *CalendarController) GetAll() {
	cm := new(managers.CalendarManagerImpl)
	events, err := cm.GetAllEvents()
	if err == nil {
		cc.Data["Events"] = events
		cc.TplName = "calendar.tpl"
	} else {
		fmt.Println(err)
		cc.handleError(err)
	}
}

// GetById is @router /calendar/:id
func (cc *CalendarController) GetById() {
	cm := new(managers.CalendarManagerImpl)
	id := cc.Ctx.Input.Param(":id")
	fmt.Println(id)
	event, err := cm.GetCalendarEventById(id)
	if err == nil {
		response := CalendarEventResponse{
			CalendarEvents: []models.CalendarEvent{event},
		}
		fmt.Println(response)
		cc.Data["json"] = &response
		cc.ServeJSON()
	} else {
		fmt.Println(err)
		cc.handleError(err)
	}
}

func (cc *CalendarController) handleError(err error) {
	switch err {
	case orm.ErrNoRows:
		fmt.Println("No result found.")
		cc.Abort(NotFound)
	case orm.ErrMissPK:
		fmt.Println("No primary key found.")
		cc.Abort(NotFound)
	default:
		fmt.Println("Internal error.")
		cc.Abort(InternalError)
	}
}
