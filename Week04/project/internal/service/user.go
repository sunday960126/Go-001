package service

import "sunday960126/Go-001/Week04/project/internal/dao"

var User UserService

type UserService interface {
}

type userService struct {
	store dao.UserStore
}

func NewUserService(store dao.UserStore) UserService {
	return &userService{
		store: store,
	}
}
