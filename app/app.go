package app

import (
	config "integrations/cmd/config"
	"integrations/logger"
	mongostorage "integrations/mongostorage"
)

// App := app struct containing resources to implement business login
type App struct {
	MongoDB *mongostorage.MongoDB
	Logger  *logger.Logger
	Config  *config.Config
}

// NewApp := returns new app object
func NewApp(c *config.Config) *App {
	m := mongostorage.NewMongoDB(&c.DatabaseConfig)
	l := logger.NewLogger(true, true, "catalogService")
	app := App{
		MongoDB: m,
		Logger:  l,
		Config:  c,
	}
	return &app
}
