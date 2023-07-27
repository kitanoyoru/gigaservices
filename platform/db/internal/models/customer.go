package models

type Customer struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
