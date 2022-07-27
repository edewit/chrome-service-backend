package database

import (
	"fmt"

	"github.com/RedHatInsights/chrome-service-backend/config"
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Init() {
  var err error
  var dialector gorm.Dialector

  cfg := config.Get()
  var dbdns string

  dbdns = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSSLMode)
		if cfg.DbSSLRootCert != "" {
			dbdns = fmt.Sprintf("%s  sslrootcert=%s", dbdns, cfg.DbSSLRootCert)
		}

		dialector = postgres.Open(dbdns)

  DB, err = gorm.Open(dialector, &gorm.Config{})

  if err != nil {
    panic(fmt.Sprintf("Database connection failed: %s", err.Error()))
  }

  fmt.Print("Database connection succesful")
}