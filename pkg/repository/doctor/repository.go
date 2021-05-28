package doctor

import (
	"errors"

	"github.com/ceng316/dentist-backend/pkg/model"
)

type Reader interface {
	GetDoctorFromID(id int64) (*model.Doctor, error)
	CheckExists(id uint) (bool, error)
}

type Writer interface {
	Add(doctor *model.Doctor) (bool, error)
	Update(doctor *model.Doctor) (bool, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

var (
	ErrSupportNotFound = errors.New("doctor not found")
)
