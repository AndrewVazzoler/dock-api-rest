package postgres

import (
	"fmt"
	"log"

	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/config"
	customer "github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/database/postgres/customer"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDB bool
	Env           string
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	fmt.Println(d.DsnTest)
	if d.Env != "test" {
		fmt.Println("aqui")
		d.Db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: d.Dsn,
		}), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.AutoMigrateDB {
		d.Db.AutoMigrate(
			&customer.Customer{},
		)
	}

	return d.Db, nil

}

func NewDb() *Database {
	return &Database{}
}

func NewDbAndConnect(cfg config.Config) *gorm.DB {
	db := NewDb()

	db.Dsn = cfg.DSN
	db.DbType = cfg.DBType
	db.DbTypeTest = cfg.DBTypeTest
	db.DsnTest = cfg.DSNTest
	db.Debug = cfg.Debug
	db.AutoMigrateDB = cfg.AutoMigrateDB
	db.Env = cfg.Env

	connection, err := db.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}
	return connection
}
