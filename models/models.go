package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type CalendarEvent struct {
	Id				string `orm:"column(id);pk"` // Set the primary key
	Description		string
	Location		string
	Attendee		string
	Time			orm.DateTimeField
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(CalendarEvent))
}
