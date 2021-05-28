package typeofappointment

import (
	"errors"

	"github.com/ceng316/dentist-backend/pkg/model"
)

type Reader interface {
	GetTypeFromID(id int64) (*model.Type, error)
	CheckExists(id uint) (bool, error)
}

type Writer interface {
	Add(t *model.Type) (bool, error)
	Update(doctor *model.Type) (bool, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

var (
	ErrSupportNotFound = errors.New("type not found")
)
