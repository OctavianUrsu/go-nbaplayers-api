package handler

import "net/http"

func getTokenFromCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	// Obtain the session token from the requests cookies
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return nil
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	return c
}
