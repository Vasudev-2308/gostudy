package models

type User struct {
	Id      int64  `json:"id:"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Subject string `json:"subject" validate:"required"`
}
