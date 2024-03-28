package models

type Event struct {
	UserId int    `json:"user_id"`
	Date   string `json:"date"`
	Name   string `json:"name"`
}

type UpdateEvent struct {
	UserId  int    `json:"user_id"`
	OldDate string `json:"old_date"`
	OldName string `json:"old_name"`
	NewDate string `json:"new_date"`
	NewName string `json:"new_name"`
}

type MonthEvents struct {
	UserId int `json:"user_id"`
	Month  int `json:"month"`
}
