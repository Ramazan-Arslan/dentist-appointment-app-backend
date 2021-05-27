package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ceng316/dentist-backend/pkg/api/response"
	"github.com/gorilla/mux"
)

// login handles login info request
func (a *API) DoctorInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["doctorID"], 10, 64)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting doctor info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	doctorInfo, err := a.service.GetDoctorService().GetDoctorInfo(id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting doctor info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, doctorInfo)
	return
}
