package appointment

import (
	"errors"

	"github.com/ceng316/dentist-backend/pkg/model"
)

type Reader interface {
	GetAll() ([]*model.Appointment, error)
}

type Writer interface {
	Add(*model.Appointment) (bool, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

var (
	ErrSupportNotFound = errors.New("appointment not found")
)
