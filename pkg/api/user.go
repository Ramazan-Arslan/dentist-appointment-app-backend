package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ceng316/dentist-backend/pkg/api/response"
	"github.com/ceng316/dentist-backend/pkg/model"
)

// login handles login info request
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	var fwReq model.User
	err := json.NewDecoder(r.Body).Decode(&fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, "")
		return
	}
	// get user info
	UserInfo, err := a.service.GetUserService().Login(fwReq)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting login info: %v", err), http.StatusBadRequest, err.Error())
		return
	}

	// write response
	response.Write(w, r, UserInfo)
	return
}
