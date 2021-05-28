package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ceng316/dentist-backend/pkg/api/response"
	"github.com/ceng316/dentist-backend/pkg/model"
	"github.com/gorilla/mux"
)

func (a *API) GetAllAppointments(w http.ResponseWriter, r *http.Request) {

	appointments, err := a.service.GetAppointmentService().GetAppointments()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting appointments : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// write response
	response.Write(w, r, appointments)
	return
}

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

func (a *API) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	var fwReq model.Appointment
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// update type
	appointmentInfo, err := a.service.GetAppointmentService().UpdateAppointment(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, appointmentInfo)
	return
}

// Delete Appointment handler
func (a *API) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["appointmentID"], 10, 64)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting doctor info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// delete type
	appointmentInfo, err := a.service.GetAppointmentService().DeleteAppointment(uint(id))
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, appointmentInfo)
	return
}
