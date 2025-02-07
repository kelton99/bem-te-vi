package http

import (
	"application"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// func dump(data interface{}) {
//     jsonData, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		fmt.Println("Erro ao serializar dados:", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(string(jsonData))
// }

// func dd(data interface{}) {
//     dump(data)
// 	os.Exit(0)
// }

type errorResponse struct {
	Error string `json:"error"`
}

// Helper functions to reduce duplication
func (s *Server) parseIDFromRequest(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.sendError(w, http.StatusBadRequest, "Invalid user ID")
		return 0, false
	}
	return id, true
}

func (s *Server) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) sendError(w http.ResponseWriter, statusCode int, message string) {
	s.sendJSONResponse(w, statusCode, errorResponse{Error: message})
}

func (s *Server) decodeJSONBody(w http.ResponseWriter, r *http.Request, dest interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		s.sendError(w, http.StatusBadRequest, "Invalid request payload")
		return false
	}
	return true
}

func (s *Server) RegisterUserRoutes(router chi.Router) {
	router.Get("/users", s.handleGetUsers)
	router.Get("/users/{id}", s.handleGetUser)
	router.Post("/users", s.handleCreateUser)
	router.Put("/users/{id}", s.handleUpdateUser)
	router.Delete("/users/{id}", s.handleDeleteUser)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, ok := s.parseIDFromRequest(w, r)
	if !ok {
		return
	}

	user, err := s.userService.GetUser(id)
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.sendJSONResponse(w, http.StatusOK, user)
}

func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userService.GetUsers()
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userDTOs := application.ToUserResponses(users) // Works with empty slice
	s.sendJSONResponse(w, http.StatusOK, userDTOs)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user application.User
	if !s.decodeJSONBody(w, r, &user) {
		return
	}

	createdUser, err := s.userService.CreateUser(user)
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.sendJSONResponse(w, http.StatusCreated, createdUser)
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := s.parseIDFromRequest(w, r)
	if !ok {
		return
	}

	var user application.User
	if !s.decodeJSONBody(w, r, &user) {
		return
	}

	user.ID = id // Ensure ID from URL takes precedence
	updatedUser, err := s.userService.UpdateUser(user)
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.sendJSONResponse(w, http.StatusOK, updatedUser)
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, ok := s.parseIDFromRequest(w, r)
	if !ok {
		return
	}

	if err := s.userService.DeleteUser(id); err != nil {
		s.sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s.sendJSONResponse(w, http.StatusNoContent, nil)
}
