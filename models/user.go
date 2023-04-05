package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"` //annotation that makes hides password from requests
	Role     Role   `json:"role"`
}

type Role int

const (
	BaseRole Role = iota
	AdminRole
)
