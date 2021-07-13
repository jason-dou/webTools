package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type CalendarEvent struct {
	Id				string `orm:"column(id);pk"` // Set the primary key
	Description		string
	Location		string
	Attendee		string
	Time			time.Time
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(CalendarEvent))
}
