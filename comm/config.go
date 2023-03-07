package comm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBX *gorm.DB

func InitDB() {
	dbx, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:12345@tcp(localhost:3306)/test?charset=utf8mb4",
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBX = dbx
}
