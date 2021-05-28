package configs

import (
	"github.com/ad3n/loyalti/controllers"
	"github.com/ad3n/loyalti/middlewares"
	"github.com/ad3n/loyalti/models"
	"github.com/ad3n/loyalti/repositories"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	connection := Connect()
	connection.AutoMigrate(&models.User{}, &models.Currency{})

	userRepository := repositories.UserRepository{Storage: connection}
	currencyRepository := repositories.CurrencyRepository{Storage: connection}

	login := controllers.Login{Repository: &userRepository}
	user := controllers.User{Repository: &userRepository}
	currency := controllers.Currency{Repository: &currencyRepository}

	demo := router.Group("/demo/v1/api")
	{
		demo.POST("/pengguna", user.Create)
		demo.PUT("/pengguna/:id", user.Update)
		demo.DELETE("/pengguna/:id", user.Delete)
		demo.GET("/pengguna/:id", user.Get)
		demo.GET("/pengguna", user.GetAll)

		demo.POST("/valuta", middlewares.ValidateToken(), currency.Create)
		demo.PUT("/valuta/:id", middlewares.ValidateToken(), currency.Update)
		demo.DELETE("/valuta/:id", middlewares.ValidateToken(), currency.Delete)
		demo.GET("/valuta/:id", middlewares.ValidateToken(), currency.Get)
		demo.GET("/valuta", middlewares.ValidateToken(), currency.GetAll)

		demo.POST("/login", login.Auth)
	}
}
