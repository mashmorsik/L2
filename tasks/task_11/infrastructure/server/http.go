package server

import (
	"encoding/json"
	"fmt"
	"github.com/mashmorsik/L2/tasks/task_11/internal/calendar"
	"github.com/mashmorsik/L2/tasks/task_11/pkg/middleware"
	"github.com/mashmorsik/L2/tasks/task_11/pkg/models"
	"net/http"
	"strconv"
)

const port = ":8080"

type HTTPServer struct {
	e calendar.Events
}

func NewServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) StartServer() error {
	http.HandleFunc("/create_event", middleware.LoggingMiddleware(s.createEventHandler))
	http.HandleFunc("/update_event", middleware.LoggingMiddleware(s.updateEventHandler))
	http.HandleFunc("/delete_event", middleware.LoggingMiddleware(s.deleteEventHandler))
	http.HandleFunc("/events_for_day", middleware.LoggingMiddleware(s.eventsForDayHandler))
	http.HandleFunc("/events_for_week", middleware.LoggingMiddleware(s.eventsForWeekHandler))
	http.HandleFunc("/events_for_month", middleware.LoggingMiddleware(s.eventsForMonthHandler))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("server can't ListenAndServe http requests: %v", err)
		return err
	}

	return nil
}

func (s *HTTPServer) createEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Printf("can't unmarshal body: %v, err: %s", r.Body, err)
		http.Error(w, "Bad Request: invalid JSON", http.StatusBadRequest)
		return
	}

	err = s.e.CreateEvent(&event)
	if err != nil {
		fmt.Printf("unable to create event: %s for user: %d", event.Name, event.UserId)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	err = s.SendResponse(w)
	if err != nil {
		fmt.Printf("unable to send response, err: %v", err)
		return
	}
}

func (s *HTTPServer) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.UpdateEvent{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Printf("can't unmarshal body: %v, err: %s", r.Body, err)
		http.Error(w, "Bad Request: invalid JSON", http.StatusBadRequest)
		return
	}

	err = s.e.UpdateEvent(&event)
	if err != nil {
		fmt.Printf("unable to update event: %s for user: %d", event.OldDate, event.UserId)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	err = s.SendResponse(w)
	if err != nil {
		fmt.Printf("unable to send response, err: %v", err)
		return
	}
}

func (s *HTTPServer) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Printf("can't unmarshal body: %v, err: %s", r.Body, err)
		http.Error(w, "Bad Request: invalid JSON", http.StatusBadRequest)
		return
	}

	err = s.e.DeleteEvent(&event)
	if err != nil {
		fmt.Printf("unable to delete event: %s for user: %d", event.Name, event.UserId)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	err = s.SendResponse(w)
	if err != nil {
		fmt.Printf("unable to send response, err: %v", err)
		return
	}
}

func (s *HTTPServer) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	idStr := r.URL.Query().Get("user_id")
	event.UserId, _ = strconv.Atoi(idStr)
	event.Date = r.URL.Query().Get("date")

	if event.UserId == 0 || event.Date == "" {
		http.Error(w, "Bad Request: Missing required parameters", http.StatusBadRequest)
		return
	}

	err, events := s.e.GetDayEvents(&event)
	if err != nil {
		fmt.Printf("unable to get day events: %s for user: %d", event.Name, event.UserId)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	for _, e := range events {
		fmt.Fprintln(w, e)
	}
}

func (s *HTTPServer) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	idStr := r.URL.Query().Get("user_id")
	event.UserId, _ = strconv.Atoi(idStr)
	event.Date = r.URL.Query().Get("date")

	if event.UserId == 0 || event.Date == "" {
		http.Error(w, "Bad Request: Missing required parameters", http.StatusBadRequest)
		return
	}

	err, events := s.e.GetWeekEvents(&event)
	if err != nil {
		fmt.Printf("unable to get week events: %s for user: %d", event.Name, event.UserId)
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	for _, e := range events {
		fmt.Fprintln(w, e)
	}
}

func (s *HTTPServer) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	event := models.MonthEvents{}

	idStr := r.URL.Query().Get("user_id")
	event.UserId, _ = strconv.Atoi(idStr)

	monthStr := r.URL.Query().Get("month")
	event.Month, _ = strconv.Atoi(monthStr)

	if event.UserId == 0 || event.Month == 0 {
		http.Error(w, "Bad Request: Missing required parameters", http.StatusBadRequest)
		return
	}

	err, events := s.e.GetMonthEvents(&event)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		err = s.SendError(w)
		if err != nil {
			fmt.Printf("unable to send err response, err: %v", err)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	for _, e := range events {
		fmt.Fprintln(w, e)
	}
}

func (s *HTTPServer) SendResponse(w http.ResponseWriter) error {
	response := models.Response{
		Result: "...",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (s *HTTPServer) SendError(w http.ResponseWriter) error {
	errResp := models.ErrorResp{
		ErrorResp: "...",
	}

	jsonData, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
