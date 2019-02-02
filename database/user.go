package database

// User type
type User struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}
