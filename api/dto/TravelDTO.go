package dto

type TravelDTO struct {
	ID          uint   `json:"id"`
	UserIDAdmin uint   `json:"user_id_admin"`
	CategoryID  uint   `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	PhotoURL    string `json:"photo_url"`
	Rating      string `json:"rating"`
	Category    string `json:"category"`
}
