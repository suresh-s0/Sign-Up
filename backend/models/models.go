package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	IsAdmin  bool   `json:"isadmin" gorm:"column:is_admin;default:false"`
	Token    string `json:"token" gorm:"column:token"`
}
