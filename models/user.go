package models

type User struct {
	ID       uint
	Username string
	Email    string `gorm:"unique"`
	Password []byte
}

type YahooUser struct {
	ID       uint
	Name     string
	Email    string
	TeamName string
}
