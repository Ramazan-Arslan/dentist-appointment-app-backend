package appointment

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

func (s *Service) GetAppointments() ([]*model.Appointment, error) {

	a, err := s.repository.GetAppointmentRepository().GetAll()
	for index, appointment := range a {
		d, err := s.repository.GetDoctorRepository().GetDoctorFromID(int64(appointment.Doctor.ID))
		if err != nil {
			continue
		}
		a[index].Doctor = d
	}
	if err != nil {
		return nil, err
	}

	return a, err
}
