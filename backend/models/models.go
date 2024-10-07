package models

type User struct {
	Id       int    `json:"id" gorm:"coloum:id"`
	Name     string `json:"name" gorm:"coloum:name"`
	Email    string `json:"email" gorm:"coloum:email"`
	Password string `json:"password" gorm:"coloum:password"`
	IsAdmin  bool   ` gorm:"coloum:isadmin"`
}
