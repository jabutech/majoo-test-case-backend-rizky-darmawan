package repository

import (
	"log"

	"github.com/majoo-test/models/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	FindByUserName(username string) (domain.User, error)
	UpdatePasswordHash()
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserName(username string) (domain.User, error) {
	var user domain.User

	err := r.db.Table("Users").Where("user_name = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdatePasswordHash() {
	passwordHash1, err := bcrypt.GenerateFromPassword([]byte("admin1"), bcrypt.MinCost)
	if err != nil {
		log.Panic(err)
	}

	err = r.db.Table("Users").Where("id = ?", 1).Update("password", string(passwordHash1)).Error
	if err != nil {
		log.Panic(err)
	}

	passwordHash2, err := bcrypt.GenerateFromPassword([]byte("admin2"), bcrypt.MinCost)
	if err != nil {
		log.Panic(err)
	}

	err = r.db.Table("Users").Where("id = ?", 2).Update("password", string(passwordHash2)).Error
	if err != nil {
		log.Panic(err)
	}
}
