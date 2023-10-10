package main

import (
	"Docker/App/config"
	database "Docker/App/databases"
	"Docker/App/migration"
	"Docker/route"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)
	migration.InitMigrationMysql(db)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routes.InitRouter(e, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
	// e.Logger.Fatal(e.Start(":80"))
}