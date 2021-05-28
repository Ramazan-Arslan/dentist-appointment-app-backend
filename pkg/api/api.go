package api

import (
	"net/http"

	"github.com/ceng316/dentist-backend/pkg/service"

	"github.com/gorilla/mux"
)

// Config represents the API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router

	config  *Config
	service service.Service
}

// New returns the api settings
func New(config *Config, svc service.Service, router *mux.Router) (*API, error) {
	api := &API{
		config:  config,
		service: svc,
		Router:  router,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// user endpoints
	api.Router.HandleFunc("/api/v1/login", api.corsMiddleware(api.logMiddleware(api.Login))).Methods("POST")
	api.Router.HandleFunc("/api/v1/userInfo", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.Login)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/doctorInfo/{doctorID}", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.DoctorInfo)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/add/doctor", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.AddDoctor)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/appointments", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.GetAllAppointments)))).Methods("GET")
	api.Router.HandleFunc("/api/v1/typeInfo/{typeID}", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.TypeInfo)))).Methods("POST")
	api.Router.HandleFunc("/api/v1/add/type", api.corsMiddleware(api.jwtMiddleware(api.logMiddleware(api.AddType)))).Methods("POST")

	return api, nil

}

// preflightHandler is the healtcheck handler
func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
