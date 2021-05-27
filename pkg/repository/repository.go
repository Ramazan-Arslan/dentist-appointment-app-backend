package repository

import (
	"github.com/ceng316/dentist-backend/pkg/repository/doctor"
	"github.com/ceng316/dentist-backend/pkg/repository/user"
)

// Repository defines the method for all operations related with repository
// Repository interface is composition of  Repository interfaces of imported packages.
type Repository interface {
	GetUserRepository() user.Repository
	GetDoctorRepository() doctor.Repository
	Shutdown()
}
