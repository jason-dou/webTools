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
	cc.loadPage()
}

// GetById is @router GET /calendar/:id
func (cc *CalendarController) GetById() {
	calendarManager := new(managers.CalendarManagerImpl)
	id := cc.Ctx.Input.Param(":id")
	fmt.Println(id)
	event, err := calendarManager.GetCalendarEventById(id)
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

// CreateCalendarEvent is @router POST /calendar
func (cc *CalendarController) CreateCalendarEvent() {
	calendarManager := new(managers.CalendarManagerImpl)
	calendarEvent := models.CalendarEvent{}
	cc.ParseForm(&calendarEvent)

	createErr := calendarManager.CreateCalendarEvent(calendarEvent)
	if createErr == nil {
		cc.loadPage()
	} else {
		fmt.Println(createErr)
		cc.handleError(createErr)
	}
}

// DeleteCalendarEvent is @router DELETE /calendar/:id
func (cc *CalendarController) DeleteCalendarEvent() {
	calendarManager := new(managers.CalendarManagerImpl)
	id := cc.Ctx.Input.Param(":id")
	deleteError := calendarManager.DeleteCalendarEvent(id)
	if deleteError != nil {
		fmt.Println(deleteError)
		cc.handleError(deleteError)
	} else {
		cc.Abort(Success)
	}
}

func (cc *CalendarController) loadPage() {
	calendarManager := new(managers.CalendarManagerImpl)
	events, err := calendarManager.GetAllEvents()
	if err == nil {
		cc.Data["Events"] = events
		cc.TplName = "calendar.tpl"
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
