package main

import (
	"desafio-criptografia/internal/handler"

	_ "desafio-criptografia/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/makeTransaction", handler.MakeTransaction)
	router.DELETE("/deleteTransaction/:id", handler.DeleteTransaction)
	router.GET("/getTransaction/:id", handler.GetTransaction)
	router.Run("localhost:8080")
}
