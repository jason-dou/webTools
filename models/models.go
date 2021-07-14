package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type CalendarEvent struct {
	Id				string `orm:"column(id);pk"` // Set the primary key
	Description		string `form:"description"`
	Location		string `form:"location"`
	Attendee		string `form:"attendee"`
	Time			time.Time `form:"time"`
	Active			bool
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(CalendarEvent))
}
