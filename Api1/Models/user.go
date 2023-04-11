package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
