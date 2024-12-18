package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Dialector = postgres.Open("host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=prefer TimeZone=Asia/Shanghai")

func ConnectToDB() {

	// url := "postgres://postgres:password@localhost:5432/postgres"
	conn, err := gorm.Open(Dialector, &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = conn
}
