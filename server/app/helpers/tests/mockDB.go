package tests

import (
	"OWallet.com/app/database"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockDB() error {
	testDB, _, err := sqlmock.New()
	if err != nil {
		return err
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})
	database.DB, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return err
	}

	return err
}
