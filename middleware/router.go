package middleware

import (
	_ "fanify/docs"
	user "fanify/internal/user/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TokenApi
// @version 0.1
// @description API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ResponseHandler())

	//Use response, but not Token
	r.GET("/token", generateTokenHandler)
	r.PUT("/login")

	auth := r.Group("/api")
	auth.Use(authMiddleware)

	//Response and token service
	auth.POST("/create-user", user.CreateUser)
	auth.POST("/create-contract")
	auth.POST("/create-token")
	auth.POST("/make-transfer")
	auth.POST("/deploy-contract")
	auth.GET("/contract-details")
	auth.GET("/all-contracts")
	auth.GET("/token")
	auth.GET("/all-token")

	auth.PUT("/logged")
	auth.GET("/user")

	return r
}
