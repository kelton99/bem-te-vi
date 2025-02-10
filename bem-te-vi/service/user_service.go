package service

import (
	"application/domain"
	"application/repository"
	"fmt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (s *UserService) CreateUser(user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("invalid user: %w", err)
	}

	err := s.UserRepository.Create(user)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (domain.User, error) {

	user, err := s.UserRepository.FindById(id)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUsers() ([]domain.User, error) {

	rows, err := s.UserRepository.FindAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Document); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) UpdateUser(user domain.User) (domain.User, error) {

	if err := user.ValidateUpdate(); err != nil {
		return domain.User{}, fmt.Errorf("invalid user: %w", err)
	}

	_, err := s.UserRepository.Update(user)

	if err != nil {
		return domain.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {

	err := s.UserRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
