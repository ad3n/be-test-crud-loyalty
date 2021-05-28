package controllers

import (
	"net/http"
	"strconv"

	"github.com/ad3n/loyalti/models"
	"github.com/ad3n/loyalti/repositories"
	"github.com/ad3n/loyalti/views"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Repository *repositories.UserRepository
}

func (u User) Create(c *gin.Context) {
	user := views.UserPost{}
	model := models.User{}

	c.BindJSON(&user)
	err := user.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})

		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	copier.Copy(&model, &user)
	model.Password = string(hash)

	err = u.Repository.Save(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terjadi kesalahan sistem",
		})

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Pengguna berhasil disimpan",
		"data": views.UserResponse{
			ID:       model.ID,
			Username: model.Username,
			Status:   model.Status,
		},
	})
}

func (u User) Update(c *gin.Context) {
	user := views.UserPut{}
	model := models.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	model.ID = id

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Pengguna tidak ditemukan",
		})

		return
	}

	c.BindJSON(&user)
	if user.Validate() != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Data tidak valid",
		})

		return
	}

	copier.Copy(&model, &user)
	model.ID = id

	err = u.Repository.Save(&model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Terjadi kesalahan sistem",
		})

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Pengguna berhasil diperbarui",
		"data": views.UserResponse{
			ID:       model.ID,
			Username: model.Username,
			Status:   model.Status,
		},
	})
}

func (u User) Delete(c *gin.Context) {
	model := models.User{}

	model.ID, _ = strconv.Atoi(c.Param("id"))

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Pengguna tidak ditemukan",
		})

		return
	}

	u.Repository.Storage.Delete(&model)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Pengguna berhasil dihapus",
	})
}

func (u User) GetAll(c *gin.Context) {
	models := []models.User{}
	users := []views.UserResponse{}

	u.Repository.All(&models)

	for _, m := range models {
		users = append(users, views.UserResponse{
			ID:       m.ID,
			Username: m.Username,
			Status:   m.Status,
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

func (u User) Get(c *gin.Context) {
	user := views.UserResponse{}
	model := models.User{}

	model.ID, _ = strconv.Atoi(c.Param("id"))

	err := u.Repository.Find(&model)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Pengguna tidak ditemukan",
		})

		return
	}

	copier.Copy(&user, &model)

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}
