package doctor

import (
	"fmt"

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

func (s *Service) UpdateDoctor(doctor model.Doctor) (bool, error) {

	// check doctor exists
	exists, err := s.repository.GetDoctorRepository().CheckExists(doctor.ID)
	if err != nil {
		fmt.Println("qsdasdasd")
		return false, err
	}
	if !exists {
		fmt.Println("bbbbb")

		return false, err
	}

	// update doctor to the repository
	boolValue, err := s.repository.GetDoctorRepository().Update(&doctor)
	if err != nil {
		fmt.Println("23232323")

		return false, nil
	}
	fmt.Println("55555")

	return boolValue, nil
}
