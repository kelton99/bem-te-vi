package interfaces

import "application/domain"

type TokenService interface {
	GenerateToken(AuthInfo *domain.AuthInfo) (*domain.Token, error)
	ValidateToken(token string) error
}
