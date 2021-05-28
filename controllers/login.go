package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ad3n/loyalti/models"
	"github.com/ad3n/loyalti/repositories"
	"github.com/ad3n/loyalti/utils"
	"github.com/ad3n/loyalti/views"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Repository *repositories.UserRepository
}

func (l Login) Auth(c *gin.Context) {
	model := models.User{}
	login := views.Login{}

	c.BindJSON(&login)

	model.Username = login.Username
	l.Repository.FindByUsername(&model)

	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Gagal Login",
		})

		return
	}

	now := time.Now()
	stringUtil := utils.Strings{}
	token := fmt.Sprintf("%d%s%s", model.ID, now.Format("1504"), stringUtil.Random(5))

	c.JSON(http.StatusOK, map[string]string{
		"message": "Sukses Login",
		"token":   token,
	})
}
