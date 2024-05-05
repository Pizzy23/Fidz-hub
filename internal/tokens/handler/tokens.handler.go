package tokens

import "github.com/gin-gonic/gin"

// @Summary Make Tokens
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.TokenCreateInput true "Data for make a new token"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.TokenCreateOutput "New Token Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to created a new token"
// @Router /api/create-token [post]
func CreateToken(c *gin.Context) {}

// @Summary Make a transfer
// @Description Create a transfer from a user to a user
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.UserController true "Data for make a new user"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.UserOutputController "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/create-user [post]
func MakeTransfer(c *gin.Context) {}

// @Summary Create user
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.UserController true "Data for make a new user"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.UserOutputController "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/create-user [post]
func GetToken(c *gin.Context) {}

// @Summary Create user
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.UserController true "Data for make a new user"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.UserOutputController "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/create-user [post]
func GetAllToken(c *gin.Context) {}

// @Summary Make a Trade
// @Description Create a transfer from a user to a new token
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.UserController true "Data for make a new user"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.UserOutputController "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/create-user [post]
func MakeTrade(c *gin.Context) {}
