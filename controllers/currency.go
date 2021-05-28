package controllers

import (
	"net/http"
	"strconv"

	"github.com/ad3n/loyalti/models"
	"github.com/ad3n/loyalti/repositories"
	"github.com/ad3n/loyalti/views"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Currency struct {
	Repository *repositories.CurrencyRepository
}

func (u Currency) Create(c *gin.Context) {
	currency := views.Currency{}
	model := models.Currency{}

	c.BindJSON(&currency)
	err := currency.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})

		return
	}

	copier.Copy(&model, &currency)

	err = u.Repository.Save(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terjadi kesalahan sistem",
		})

		return
	}

	currency.ID = model.ID

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Valuta berhasil disimpan",
		"data":    currency,
	})
}

func (u Currency) Update(c *gin.Context) {
	currency := views.Currency{}
	model := models.Currency{}

	model.ID, _ = strconv.Atoi(c.Param("id"))

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Valuta tidak ditemukan",
		})

		return
	}

	c.BindJSON(&currency)
	if currency.Validate() != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Data tidak valid",
		})

		return
	}

	currency.ID = model.ID

	copier.Copy(&model, &currency)

	err = u.Repository.Save(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terjadi kesalahan sistem",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Valuta berhasil diperbarui",
		"data":    currency,
	})
}

func (u Currency) Delete(c *gin.Context) {
	model := models.Currency{}

	model.ID, _ = strconv.Atoi(c.Param("id"))

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Valuta tidak ditemukan",
		})

		return
	}

	u.Repository.Storage.Delete(&model)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Valuta berhasil dihapus",
	})
}

func (u Currency) GetAll(c *gin.Context) {
	models := []models.Currency{}
	Currencys := []views.Currency{}

	u.Repository.All(&models)

	for _, m := range models {
		Currencys = append(Currencys, views.Currency{
			ID:    m.ID,
			Code:  m.Code,
			Name:  m.Name,
			Price: m.Price,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": Currencys,
	})
}

func (u Currency) Get(c *gin.Context) {
	currency := views.Currency{}
	model := models.Currency{}

	model.ID, _ = strconv.Atoi(c.Param("id"))

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Valuta tidak ditemukan",
		})

		return
	}

	copier.Copy(&currency, &model)

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": currency,
	})
}
