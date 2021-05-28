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

// Doctor info handler
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

// Add doctor handler
func (a *API) AddDoctor(w http.ResponseWriter, r *http.Request) {
	var fwReq model.Doctor
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// get user info
	doctorInfo, err := a.service.GetDoctorService().AddDoctor(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, doctorInfo)
	return
}

// Update doctor handler
func (a *API) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	var fwReq model.Doctor
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// get user info
	doctorInfo, err := a.service.GetDoctorService().UpdateDoctor(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, doctorInfo)
	return
}
