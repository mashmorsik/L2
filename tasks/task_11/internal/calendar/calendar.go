package calendar

import (
	"errors"
	"fmt"
	"github.com/mashmorsik/L2/tasks/task_11/pkg/models"
	"time"
)

type Events struct {
	events map[int]map[time.Time][]string
}

func NewEvents() *Events {
	return &Events{events: make(map[int]map[time.Time][]string)}
}

func (e *Events) CreateEvent(event *models.Event) error {
	if e.events == nil {
		e.events = make(map[int]map[time.Time][]string)
	}

	dateFmt, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return err
	}

	innerMap, ok := e.events[event.UserId]
	if !ok {
		innerMap = make(map[time.Time][]string)
		e.events[event.UserId] = innerMap
	}

	innerMap[dateFmt] = append(innerMap[dateFmt], event.Name)

	return nil
}

func (e *Events) UpdateEvent(event *models.UpdateEvent) error {
	existingDate, err := time.Parse("2006-01-02", event.OldDate)
	if err != nil {
		return fmt.Errorf("error parsing existing date: %v", err)
	}

	dateToUpdate, err := time.Parse("2006-01-02", event.NewDate)
	if err != nil {
		return fmt.Errorf("error parsing new date: %v", err)
	}

	userEvents, ok := e.events[event.UserId]
	if !ok {
		return fmt.Errorf("no events found for user %d", event.UserId)
	}

	events, ok := userEvents[existingDate]
	if !ok {
		return fmt.Errorf("no event found for user %d on date %s", event.UserId, event.UserId)
	}

	var eventIndex = -1
	for i, ev := range events {
		if ev == event.OldName {
			eventIndex = i
			break
		}
	}

	if eventIndex != -1 {
		events[eventIndex] = event.NewName
		delete(userEvents, existingDate)
		userEvents[dateToUpdate] = events
		e.events[event.UserId] = userEvents
	} else {
		return fmt.Errorf("no event '%s' found for user %d on date %s",
			event.OldName, event.UserId, event.OldDate)
	}

	return nil
}

func (e *Events) DeleteEvent(event *models.Event) error {
	dateFmt, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return fmt.Errorf("error parsing existing date: %v", err)
	}

	userEvents, ok := e.events[event.UserId]
	if !ok {
		return fmt.Errorf("no events found for user %d", event.UserId)
	}

	events, ok := userEvents[dateFmt]
	if !ok {
		return fmt.Errorf("no events found for user %d on date %s", event.UserId, dateFmt)
	}

	var found bool
	for i, ev := range events {
		if ev == event.Name {
			events = append(events[:i], events[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("event '%s' not found for user %d on date %s", event, event.UserId, dateFmt)
	}

	userEvents[dateFmt] = events

	return nil
}

func (e *Events) GetDayEvents(event *models.Event) (error, []string) {
	dateFmt, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return err, nil
	}

	userEvents, ok := e.events[event.UserId]
	if !ok {
		return fmt.Errorf("no events found for user %d", event.UserId), nil
	}

	events, ok := userEvents[dateFmt]
	if !ok {
		return fmt.Errorf("no events found for user %d on date %s", event.UserId, dateFmt), nil
	}

	return nil, events
}

func (e *Events) GetWeekEvents(event *models.Event) (error, []string) {
	dateFmt, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return errors.New("invalid dateFmt format"), nil
	}

	startOfWeek := dateFmt.AddDate(0, 0, -int(dateFmt.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	var weekEvents []string

	for eventDate, events := range e.events[event.UserId] {
		if eventDate.After(startOfWeek) && eventDate.Before(endOfWeek) || eventDate.Equal(startOfWeek) || eventDate.Equal(endOfWeek) {
			weekEvents = append(weekEvents, events...)
		}
	}

	if len(weekEvents) == 0 {
		return errors.New("no events found for the specified week"), nil
	}

	return nil, weekEvents
}

func (e *Events) GetMonthEvents(event *models.MonthEvents) (error, []string) {
	var monthEvents []string

	startOfMonth := time.Date(time.Now().Year(), time.Month(event.Month), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	for date, events := range e.events[event.UserId] {
		if date.After(startOfMonth) && date.Before(endOfMonth) || date.Equal(startOfMonth) || date.Equal(endOfMonth) {
			monthEvents = append(monthEvents, events...)
		}
	}

	if len(monthEvents) == 0 {
		return errors.New("no events found for the specified month"), nil
	}

	return nil, monthEvents
}
