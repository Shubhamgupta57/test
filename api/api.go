package api

import (
	"integrations/app"
	config "integrations/cmd/config"
	"integrations/logger"
	"integrations/validation"

	"github.com/gorilla/mux"
)

// API := returns API struct
type API struct {
	App        *app.App
	Router     *Router
	MainRouter *mux.Router
	Validator  *validation.Validation
	Logger     *logger.Logger
	Config     *config.Config
}

// NewAPI := retuns router object
func NewAPI(m *mux.Router, c *config.Config) *API {
	v := validation.NewValidation()
	l := logger.NewLogger(true, true, "api")
	api := API{
		MainRouter: m,
		Router:     &Router{},
		Validator:  v,
		Logger:     l,
		Config:     c,
	}
	api.App = app.NewApp(c)
	api.setupRoutes()
	config := config.GetConfig()
	api.Config = config
	return &api
}
