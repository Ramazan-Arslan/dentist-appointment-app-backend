package doctor

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

func (s *Service) GetDoctorInfo(id int64) (*model.Doctor, error) {

	d, err := s.repository.GetDoctorRepository().GetDoctorFromID(id)
	if err != nil {
		return nil, err
	}

	return d, err
}

func (s *Service) AddDoctor(doctor model.Doctor) (bool, error) {

	boolValue, err := s.repository.GetDoctorRepository().Add(&doctor)
	if err != nil {
		return false, err
	}

	return boolValue, err
}
