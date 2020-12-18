// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"sunday960126/Go-001/Week04/project/internal/dao"
)

func InitUserServer(db *gorm.DB) UserService {
	wire.Build(dao.NewUserStore, NewUserService)
	return User
}
