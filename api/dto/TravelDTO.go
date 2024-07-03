package dto

import "time"

type TravelDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	PhotoURL    string    `json:"photo_url"`
	Rating      string    `json:"rating"`
	Category    string    `json:"category"`
	User        string    `json:"user_nome"`
	UserPhoto   string    `json:"user_photo"`
}
