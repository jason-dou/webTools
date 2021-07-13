package queryEngines

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"webTools/models"
)

type Calendar interface {
	GetAllCalendarEvents() ([]models.CalendarEvent, error)
	GetCalendarEventById(id string) (models.CalendarEvent, error)
}

type CalendarQueryEngine struct {
	Ormer orm.Ormer
}

func (queryEngine CalendarQueryEngine) GetAllCalendarEvents() ([]models.CalendarEvent, error) {
	var events []models.CalendarEvent
	_, err := queryEngine.Ormer.QueryTable("calendar_event").All(&events)

	return events, err
}

func (queryEngine CalendarQueryEngine) GetCalendarEventById(id string) (models.CalendarEvent, error) {
	event := &models.CalendarEvent{Id: id}
	err := queryEngine.Ormer.Read(event)
	fmt.Println(event.Time)
	return *event, err
}
