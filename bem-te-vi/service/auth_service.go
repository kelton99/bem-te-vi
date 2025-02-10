package service

import (
	"application/domain"
	"application/interfaces"
	"context"
	"fmt"
	"strconv"
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

type AuthService struct {
	userService interfaces.UserService
}

func NewAuthService(userService interfaces.UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (s *AuthService) Login(ctx context.Context, query *domain.LoginUserQuery) (*domain.AuthInfo, error) {

	user, err := s.userService.GetUserByEmail(query.Username)

	if err != nil {
		return nil, err
	}

	if user.Password != query.Password {
		return nil, fmt.Errorf("invalid password")
	}

	return &domain.AuthInfo{
		AuthId: strconv.Itoa(user.ID),
		Email:  user.Email,
	}, nil
}
