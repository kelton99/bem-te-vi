package interfaces

import (
	"application/domain"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, query *domain.LoginUserQuery) (*domain.AuthInfo, error)
	// Logout(ctx context.Context, authId string) error
}
