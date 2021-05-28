package repositories

import (
	"github.com/ad3n/loyalti/models"
	"gorm.io/gorm"
)

type CurrencyRepository struct {
	Storage *gorm.DB
}

func (r *CurrencyRepository) Save(currency *models.Currency) error {
	return r.Storage.Save(currency).Error
}

func (r *CurrencyRepository) Find(currency *models.Currency) error {
	return r.Storage.First(currency).Error
}

func (r *CurrencyRepository) All(currency *[]models.Currency) error {
	return r.Storage.Find(currency).Error
}

func (r *CurrencyRepository) Remove(currency *models.Currency) error {
	return r.Storage.Save(&currency).Error
}
