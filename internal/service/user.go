package service

import (
	"user-registration/internal/model"
)

type UserService struct {
	users []model.User
}

func NewUserService() *UserService {
	return &UserService{users: []model.User{}}
}

func (s *UserService) RegisterUser(user model.User) model.User {
	s.users = append(s.users, user)
	return user
}
