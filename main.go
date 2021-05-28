package main

import (
	"fmt"

	"github.com/ad3n/loyalti/configs"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
}

func main() {
	r := gin.Default()

	configs.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%d", configs.Env.AppPort))
}
