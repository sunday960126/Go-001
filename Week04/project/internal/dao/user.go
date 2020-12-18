package dao

import "github.com/jinzhu/gorm"

type UserStore interface {
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{
		db: db,
	}
}
