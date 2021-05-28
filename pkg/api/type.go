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

// get type  info
func (a *API) TypeInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["typeID"], 10, 64)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting type info: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	typeInfo, err := a.service.GetTypeService().GetTypeInfo(id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting type info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, typeInfo)
	return
}

// login handles login info request
func (a *API) AddType(w http.ResponseWriter, r *http.Request) {
	var fwReq model.Type
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding type : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// add type info
	typeInfo, err := a.service.GetTypeService().AddType(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error adding type : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, typeInfo)
	return
}

// Update type handler
func (a *API) UpdateType(w http.ResponseWriter, r *http.Request) {
	var fwReq model.Type
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}
	// udapte type
	typeInfo, err := a.service.GetTypeService().UpdateType(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error update doctor : %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, typeInfo)
	return
}
