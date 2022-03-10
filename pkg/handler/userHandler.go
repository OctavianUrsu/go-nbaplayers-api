package handler

import (
	"encoding/json"
	"io"
	"net/http"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
	"github.com/sirupsen/logrus"
)

func (h *Handler) userSignup(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.New().Warn(err)
		return
	}

	// Store the user signup details in a Data Transfer Object
	var userSignupDTO structure.User

	json.Unmarshal(req, &userSignupDTO)

	validateErr := validate.Struct(userSignupDTO)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validateErr.Error()))
		return
	}

	if err := h.service.SignUp(userSignupDTO); err != nil {
		// resp: In case of error, write the error + the http status
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// resp: In case of success, write the created player + the http status
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "user created")
}
