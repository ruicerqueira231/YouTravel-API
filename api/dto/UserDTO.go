package dto

type UserDTO struct {
	ID       uint   `json:"id"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Photo    string `json:"photo,omitempty"` // Include only if you use photos
}
