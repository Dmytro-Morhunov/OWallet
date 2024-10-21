package database

import (
	"testing"

	"OWallet.com/database"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
)

func TestConnectToDB(t *testing.T) {
	testDB, _, _ := sqlmock.New()

	database.Dialector = postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})
	database.ConnectToDB()
	assert.NotNil(t, database.DB)
}
