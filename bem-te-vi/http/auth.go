package http

import (
	"application"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func (s *Server) RegisterAuthRoutes(router chi.Router) {
	router.Post("/login", s.handleLogin)
}


func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {

    var loginUserQuery application.LoginUserQuery

    if err := json.NewDecoder(r.Body).Decode(&loginUserQuery); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(errorResponse{Error: "Invalid request payload"})
        return
    }

	// login user
	userInfo, err := s.authService.Login(r.Context(), &loginUserQuery)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
        return
    }

    // generate token
	token, err := s.tokenService.GenerateToken(userInfo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})
		return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "token": token.Token,
        "expiry": token.Expiry,
        "email": userInfo.Email,
        "auth_id": userInfo.AuthId,		
    })
}