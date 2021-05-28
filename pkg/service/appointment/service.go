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
		t, err := s.repository.GetTypeRepository().GetTypeFromID(int64(appointment.Type.ID))
		if err != nil {
			continue
		}
		a[index].Type = t

	}
	if err != nil {
		return nil, err
	}

	return a, err
}

func (s *Service) AddAppointment(appointment model.Appointment) (bool, error) {

	boolValue, err := s.repository.GetAppointmentRepository().Add(&appointment)
	if err != nil {
		return false, err
	}

	return boolValue, err
}

func (s *Service) UpdateAppointment(appointment model.Appointment) (bool, error) {

	// check appointment exists
	exists, err := s.repository.GetAppointmentRepository().CheckExists(appointment.ID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, err
	}

	// update appointment
	boolValue, err := s.repository.GetAppointmentRepository().Update(&appointment)
	if err != nil {
		return false, nil
	}
	return boolValue, nil
}

func (s *Service) DeleteAppointment(id uint) (bool, error) {

	// check appointment exists
	exists, err := s.repository.GetAppointmentRepository().CheckExists(id)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, err
	}

	// delete appointment
	boolValue, err := s.repository.GetAppointmentRepository().Delete(uint(id))
	if err != nil {
		return false, nil
	}
	return boolValue, nil
}
