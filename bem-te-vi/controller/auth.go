package controller

import (
	"application/domain"
	"application/interfaces"
	"application/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthController struct {
	authService  interfaces.AuthService
	tokenService interfaces.TokenService
}

func NewAuthController(router chi.Router, as interfaces.AuthService, ts interfaces.TokenService) *AuthController {
	controller := &AuthController{
		authService: as,
	}

	router.Post("/login", controller.handleLogin)

	return controller
}

func (ac *AuthController) handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUserQuery domain.LoginUserQuery

	if err := json.NewDecoder(r.Body).Decode(&loginUserQuery); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	// login user
	userInfo, err := ac.authService.Login(r.Context(), &loginUserQuery)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Error: err.Error()})
		return
	}

	// generate token
	token, err := ac.tokenService.GenerateToken(userInfo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":   token.Token,
		"expiry":  token.Expiry,
		"email":   userInfo.Email,
		"auth_id": userInfo.AuthId,
	})
}
