package token

import (
	"application"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("r5g1er65g1er65g1e6r5g1er65g1e6r51ge6r51ge65rg1er651g65er1") // chave secreta usada para assinar o token

type TokenService struct {}

func NewUserService() *TokenService {
    return &TokenService{}
}

func (s *TokenService) GenerateToken(authInfo *application.AuthInfo) (*application.Token, error) {
	// Definindo as claims (dados armazenados no token)
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     authInfo.AuthId,
		"email":      authInfo.Email,
		"exp":        time.Now().Add(time.Minute * 15).Unix(), // Expira em 15 minutos
	}

	// Criando o token com o método de assinatura HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Gerando a string do token usando a chave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &application.Token{Token: tokenString, Expiry: time.Now().Add(time.Minute * 15)}, nil
}

func (s *TokenService) ValidateToken(tokenString string) error {
    // Valida se o token está no formato esperado
    if len(strings.Split(tokenString, ".")) != 3 {
        return fmt.Errorf("token no formato inválido")
    }

    // Faz o parse do token e validação da assinatura
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return fmt.Errorf("erro ao fazer parse do token: %v", err)
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return fmt.Errorf("token inválido")
    }

    expClaim, ok := claims["exp"].(float64)
    if !ok {
        return fmt.Errorf("token sem claim exp")
    }

    expiry := time.Unix(int64(expClaim), 0)
    if time.Now().After(expiry) {
        return fmt.Errorf("token expirado")
    }

    return nil // Sucesso
}

