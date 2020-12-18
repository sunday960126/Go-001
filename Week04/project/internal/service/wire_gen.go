// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

import (
	"github.com/jinzhu/gorm"
	"sunday960126/Go-001/Week04/project/internal/dao"
)

// Injectors from wire.go:

func InitUserServer(db *gorm.DB) UserService {
	userStore := dao.NewUserStore(db)
	serviceUserService := NewUserService(userStore)
	return serviceUserService
}
