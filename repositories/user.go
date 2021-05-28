package repositories

import (
	"github.com/ad3n/loyalti/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Storage *gorm.DB
}

func (r *UserRepository) Save(user *models.User) error {
	return r.Storage.Save(user).Error
}

func (r *UserRepository) Find(user *models.User) error {
	return r.Storage.First(user).Error
}

func (r *UserRepository) FindByUsername(user *models.User) error {
	return r.Storage.Where("puser = ? AND pstatus = 1", user.Username).First(user).Error
}

func (r *UserRepository) All(user *[]models.User) error {
	return r.Storage.Find(user).Error
}

func (r *UserRepository) Remove(user *models.User) error {
	return r.Storage.Save(&user).Error
}
