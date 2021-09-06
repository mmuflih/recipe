package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/recipe/api/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

type PgSQL struct {
	MainDB *gorm.DB
}

func NewPostgreSqlConn(cfg conf.Config) (*PgSQL, error) {
	app.Logger("Initial PgSql Connection")
	dbUser := cfg.GetString(`pgsql.user`)
	dbPass := cfg.GetString(`pgsql.pass`)
	dbName := cfg.GetString(`pgsql.database`)
	dbHost := cfg.GetString(`pgsql.address`)
	dbPort := cfg.GetString(`pgsql.port`)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort)
	fmt.Println("DSN:", dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	app.Logger("GORM DSN ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		app.Error(err)
		return nil, err
	}

	app.Success("PostgreSQL Connected")

	return &PgSQL{
		MainDB: db,
	}, nil
}
