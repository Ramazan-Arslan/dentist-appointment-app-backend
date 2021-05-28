package service

import (
	"github.com/ceng316/dentist-backend/pkg/service/appointment"
	"github.com/ceng316/dentist-backend/pkg/service/doctor"
	typeofappointment "github.com/ceng316/dentist-backend/pkg/service/typeOfAppointment"
	"github.com/ceng316/dentist-backend/pkg/service/user"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	GetDoctorService() *doctor.Service
	GetAppointmentService() *appointment.Service
	GetTypeService() *typeofappointment.Service
	Shutdown()
}
