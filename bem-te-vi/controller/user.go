package controller

import (
	"application/domain"
	"application/interfaces"
	"application/utils"
	"net/http"

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

type UserController struct {
	userService interfaces.UserService
}

func NewUserController(router chi.Router, us interfaces.UserService) *UserController {
	controller := &UserController{
		userService: us,
	}
	router.Get("/users", controller.handleGetUsers)
	router.Get("/users/{id}", controller.handleGetUser)
	router.Post("/users", controller.handleCreateUser)
	router.Put("/users/{id}", controller.handleUpdateUser)
	router.Delete("/users/{id}", controller.handleDeleteUser)

	return controller
}

func (s *UserController) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(w, r)
	if !ok {
		return
	}

	user, err := s.userService.GetUser(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, user)
}

func (s *UserController) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userService.GetUsers()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userDTOs := domain.ToUserResponses(users) // Works with empty slice
	utils.SendJSONResponse(w, http.StatusOK, userDTOs)
}

func (s *UserController) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if !utils.DecodeJSONBody(w, r, &user) {
		return
	}

	createdUser, err := s.userService.CreateUser(user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, createdUser)
}

func (s *UserController) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(w, r)
	if !ok {
		return
	}

	var user domain.User
	if !utils.DecodeJSONBody(w, r, &user) {
		return
	}

	user.ID = id // Ensure ID from URL takes precedence
	updatedUser, err := s.userService.UpdateUser(user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, updatedUser)
}

func (s *UserController) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.ParseIDFromRequest(w, r)
	if !ok {
		return
	}

	if err := s.userService.DeleteUser(id); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusNoContent, nil)
}
