package managers

import (
	"github.com/beego/beego/v2/client/orm"
	"webTools/models"
	"webTools/queryEngines"
)

type CalendarManager interface {
	GetAllEvents() ([]models.CalendarEvent, error)
	GetCalendarEventById(id string) (models.CalendarEvent, error)
	CreateCalendarEvent(event models.CalendarEvent) error
	DeleteCalendarEvent(id string) error
}

type CalendarManagerImpl struct {

}

func (cm CalendarManagerImpl) GetAllEvents() ([]models.CalendarEvent, error) {
	cqe := queryEngines.CalendarQueryEngine{Ormer: orm.NewOrm()}
	events, err := cqe.GetAllCalendarEvents()
	if err == nil {
		return events, nil
	}
	return []models.CalendarEvent{}, err
}

func (cm CalendarManagerImpl) GetCalendarEventById(id string) (models.CalendarEvent, error) {
	cqe := queryEngines.CalendarQueryEngine{Ormer: orm.NewOrm()}
	event, err := cqe.GetCalendarEventById(id)
	if err == nil {
		return event, nil
	}
	return models.CalendarEvent{}, err
}

func (cm CalendarManagerImpl) CreateCalendarEvent(event models.CalendarEvent) error {
	cqe := queryEngines.CalendarQueryEngine{Ormer: orm.NewOrm()}
	err := cqe.CreateCalendarEvent(event)

	return err
}

func (cm CalendarManagerImpl) DeleteCalendarEvent(id string) error {
	cqe := queryEngines.CalendarQueryEngine{Ormer: orm.NewOrm()}
	err := cqe.DeleteCalendarEvent(id)

	return err
}
