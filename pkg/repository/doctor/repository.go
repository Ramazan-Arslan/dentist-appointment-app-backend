package doctor

import (
	"errors"

	"github.com/ceng316/dentist-backend/pkg/model"
)

type Reader interface {
	GetDoctorFromID(id int64) (*model.Doctor, error)
}

type Writer interface {
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

var (
	ErrSupportNotFound = errors.New("doctor not found")
)
