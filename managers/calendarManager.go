package managers

import (
	"github.com/beego/beego/v2/client/orm"
	"webTools/models"
	"webTools/queryEngines"
)

type CalendarManager interface {
	GetAllEvents() ([]models.CalendarEvent, error)
	GetCalendarEventById(id string) (models.CalendarEvent, error)
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
