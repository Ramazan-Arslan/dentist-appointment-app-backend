package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ceng316/dentist-backend/pkg/api/response"
	"github.com/ceng316/dentist-backend/pkg/model"
)

// login handles login info request
func (a *API) GetAllAppointments(w http.ResponseWriter, r *http.Request) {

	appointments, err := a.service.GetAppointmentService().GetAppointments()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting appointments : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(appointments)
	// write response
	response.Write(w, r, appointments)
	return
}

// login handles login info request
func (a *API) AddAppointments(w http.ResponseWriter, r *http.Request) {

	var fwReq model.Appointment
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// get user info
	appointmentInfo, err := a.service.GetAppointmentService().AddAppointment(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, appointmentInfo)

	return
}
