package database

import (
	"database/sql"
	"fmt"
	"github.com/giwiro/escrap/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var (
		host     = config.Conf.Database.Host
		user     = config.Conf.Database.User
		port     = config.Conf.Database.Port
		password = config.Conf.Database.Password
		name     = config.Conf.Database.DBName
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	sqlDB, sqlDBErr := sql.Open("pgx", dsn)

	if sqlDBErr != nil {
		log.Fatalln(sqlDBErr)
	}

	db, gormDbErr := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if gormDbErr != nil {
		log.Fatalln(gormDbErr)
	}

	return db
}
