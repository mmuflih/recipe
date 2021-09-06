package provider

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/recipe/api/app"
	"github.com/mmuflih/recipe/api/app/helper"
	"github.com/mmuflih/recipe/api/config"
	"github.com/mmuflih/recipe/api/http/core/request"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func BuildConfigProvider(c *dig.Container) *dig.Container {
	/** inject env.json as config */
	if err := c.Provide(func() conf.Config {
		return conf.NewConfig()
	}); err != nil {
		panic(err)
	}

	/** inject database config */
	if err := c.Provide(func(c conf.Config) (*config.PgSQL, error) {
		debug := false
		args := os.Args
		for _, arg := range args {
			if arg == "--debug-sql-on" {
				debug = true
				break
			}
			if arg == "--debug-sql-off" {
				debug = true
				break
			}
		}
		db, err := config.NewPostgreSqlConn(c)

		if debug {
			app.Logger("SQL-DEBUGER RUNNING")
			db.MainDB = db.MainDB.Debug()
		}

		return db, err
	}); err != nil {
		panic(err)
	}

	/** inject request reader */
	if err := c.Provide(func() request.Reader {
		return request.NewRequestReader()
	}); err != nil {
		panic(err)
	}

	/** inject slug generator */
	if err := c.Provide(func(conn *config.PgSQL) helper.Slug {
		return helper.NewSlug(conn)
	}); err != nil {
		panic(err)
	}

	/** inject gin engine */
	if err := c.Provide(func(c conf.Config) *gin.Engine {
		ginMode := c.GetString("env")
		gin.SetMode(ginMode)
		router := gin.Default()
		router.Use(cors.New(cors.Config{
			AllowOrigins:     c.GetStringSlice("cors.allowed_origins"),
			AllowMethods:     c.GetStringSlice("cors.allowed_methods"),
			AllowHeaders:     c.GetStringSlice("cors.allowed_headers"),
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))
		return router
	}); err != nil {
		panic(err)
	}

	return c
}
