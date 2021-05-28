package service

import (
	"github.com/ceng316/dentist-backend/pkg/repository"
	"github.com/ceng316/dentist-backend/pkg/service/appointment"
	"github.com/ceng316/dentist-backend/pkg/service/doctor"
	typeofappointment "github.com/ceng316/dentist-backend/pkg/service/typeOfAppointment"
	"github.com/ceng316/dentist-backend/pkg/service/user"
)

type Provider struct {
	cfg *Config

	repository         repository.Repository
	userService        *user.Service
	doctorService      *doctor.Service
	appointmentService *appointment.Service
	typeService        *typeofappointment.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {

	userService, err := user.NewService(repo)
	if err != nil {
		return nil, err
	}
	doctorService, err := doctor.NewService(repo)
	if err != nil {
		return nil, err
	}
	appointmentService, err := appointment.NewService(repo)
	if err != nil {
		return nil, err
	}
	typeService, err := typeofappointment.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:                cfg,
		repository:         repo,
		userService:        userService,
		doctorService:      doctorService,
		appointmentService: appointmentService,
		typeService:        typeService,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}
func (p *Provider) GetUserService() *user.Service {
	return p.userService
}
func (p *Provider) GetDoctorService() *doctor.Service {
	return p.doctorService
}
func (p *Provider) GetAppointmentService() *appointment.Service {
	return p.appointmentService
}
func (p *Provider) GetTypeService() *typeofappointment.Service {
	return p.typeService
}
func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}
