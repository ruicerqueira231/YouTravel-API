package dto

type UserDTO struct {
	ID       uint   `json:"id"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo,omitempty"`
}
