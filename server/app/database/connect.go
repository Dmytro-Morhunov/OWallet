package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Dialector = postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5431 sslmode=disable TimeZone=Asia/Shanghai")

func ConnectToDB() {

	// url := "postgres://postgres:password@localhost:5431/postgres"
	conn, err := gorm.Open(Dialector, &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = conn
}
