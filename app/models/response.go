package models

type NewsResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PhoneNumber int    `json:"phone_number"`
}
