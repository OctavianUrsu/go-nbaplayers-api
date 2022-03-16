package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/OctavianUrsu/go-nbaplayers-api/pkg/models"
)

// Sign up handler
func (h *Handler) userSignup(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Store the user signup details in a Data Transfer Object
	var userSignupDTO models.User

	// Unmarshal the request body to variable
	if err := json.Unmarshal(req, &userSignupDTO); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Validate the request body to concise to User models
	validateErr := validate.Struct(userSignupDTO)
	if validateErr != nil {
		http.Error(w, validateErr.Error(), 400)
		return
	}

	// Sign up and get a auth token
	err = h.service.SignUp(userSignupDTO)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	} else {
		// If the user is created, response with a success message
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("user created"))
	}
}

// Sign in handler
func (h *Handler) userSignin(w http.ResponseWriter, r *http.Request) {
	// Read the request
	req, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Store the user signin details in a Data Transfer Object
	var userSigninDTO models.UserSignin

	// Unmarshal the request body to variable
	if err := json.Unmarshal(req, &userSigninDTO); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Sign in and get an auth token
	signedToken, err := h.service.SignIn(userSigninDTO)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	} else {
		// We set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		expirationTime := time.Now().Add(24 * time.Hour)

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   signedToken,
			Expires: expirationTime,
		})

		// If the user is logged in, response with a success message
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("logged in"))
	}
}
