package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DNS = "root:Addlink123!@tcp(127.0.0.1:3306)/gotut?charset=utf8&parseTime=True"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", DNS)
	return db, err
}
