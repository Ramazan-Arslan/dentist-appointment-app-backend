package typeofappointment

import (
	"github.com/ceng316/dentist-backend/pkg/model"
	"github.com/ceng316/dentist-backend/pkg/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) GetTypeInfo(id int64) (*model.Type, error) {

	d, err := s.repository.GetTypeRepository().GetTypeFromID(id)
	if err != nil {
		return nil, err
	}

	return d, err
}

func (s *Service) AddType(t model.Type) (bool, error) {

	boolValue, err := s.repository.GetTypeRepository().Add(&t)
	if err != nil {
		return false, err
	}

	return boolValue, err
}

func (s *Service) UpdateType(t model.Type) (bool, error) {

	// check type exists
	exists, err := s.repository.GetDoctorRepository().CheckExists(t.ID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, err
	}

	// update type
	boolValue, err := s.repository.GetTypeRepository().Update(&t)
	if err != nil {
		return false, nil
	}
	return boolValue, nil
}
