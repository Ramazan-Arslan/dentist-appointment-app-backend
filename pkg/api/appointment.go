package api

import (
	"fmt"
	"net/http"

	"github.com/ceng316/dentist-backend/pkg/api/response"
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
