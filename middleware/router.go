package middleware

import (
	_ "fanify/docs"
	tokens "fanify/internal/tokens/handler"
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
	r.PUT("/login", LoginUser)

	auth := r.Group("/api")
	auth.Use(authMiddleware)

	//Response and token service
	auth.POST("/create-user", user.CreateUser)
	auth.POST("/create-contract", user.CreateContract)

	auth.POST("/deploy-contract", user.DeployContract)
	auth.GET("/contract-details", user.PullContract)
	auth.GET("/all-contracts", user.PullAllContract)

	auth.PUT("/logged")
	auth.GET("/user", user.PullUser)

	auth.POST("/create-token", tokens.CreateToken)
	auth.POST("/make-transfer", tokens.MakeTransfer)
	auth.POST("/make-trade", tokens.MakeTrade)
	auth.GET("/get-token", tokens.GetToken)
	auth.GET("/all-token", tokens.GetAllToken)

	return r
}
