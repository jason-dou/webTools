package queryEngines

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/google/uuid"
	"webTools/models"
)

type Calendar interface {
	GetAllCalendarEvents() ([]models.CalendarEvent, error)
	GetCalendarEventById(id string) (models.CalendarEvent, error)
	CreateCalendarEvent(event models.CalendarEvent) error
	DeleteCalendarEvent(id string) error
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

	return *event, err
}

func (queryEngine CalendarQueryEngine) CreateCalendarEvent(event models.CalendarEvent) error {
	event.Id = uuid.NewString()
	_, err := queryEngine.Ormer.Insert(&event)

	return err
}

func (queryEngine CalendarQueryEngine) DeleteCalendarEvent(id string) error {
	_, err := queryEngine.Ormer.Delete(&models.CalendarEvent{Id: id})

	return err
}
