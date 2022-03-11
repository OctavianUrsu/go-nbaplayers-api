package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	structure "github.com/OctavianUrsu/go-nbaplayers-api"
)

// Sign up handler
func (h *Handler) userSignup(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Store the user signup details in a Data Transfer Object
	var userSignupDTO structure.User

	// Unmarshal the request body to variable
	if err := json.Unmarshal(req, &userSignupDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Validate the request body to concise to User structure
	validateErr := validate.Struct(userSignupDTO)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validateErr.Error()))
		return
	}

	// Sign up and get a auth token
	signedToken, err := h.service.SignUp(userSignupDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		// We set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		expirationTime := time.Now().Add(24 * time.Hour)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   *signedToken,
			Expires: expirationTime,
		})

		// If the user is created, response with a success message
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, "user created")
	}
}

// Sign in handler
func (h *Handler) userSignin(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Store the user signin details in a Data Transfer Object
	var userSigninDTO structure.UserSignin

	// Unmarshal the request body to variable
	if err := json.Unmarshal(req, &userSigninDTO); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Sign in and get an auth token
	signedToken, err := h.service.SignIn(userSigninDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		// We set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		expirationTime := time.Now().Add(24 * time.Hour)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   *signedToken,
			Expires: expirationTime,
		})

		// If the user is logged in, response with a success message
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "logged in")
	}
}
