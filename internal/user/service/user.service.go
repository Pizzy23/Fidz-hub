package service

import (
	"fanify/db"
	inter "fanify/internal/user/interface"
	"fanify/internal/user/storage"
	lumx "fanify/package/lumx"
	"fanify/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullContract(c *gin.Context, ContractId string) {
	if ContractId == "" {
		c.Set("Response", "Insert a ContractId")
		c.Status(http.StatusNotAcceptable)
	}
	res, err := lumx.FetchContractDetails(ContractId)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func PullAllContract(c *gin.Context) {
	res, err := lumx.FetchContracts()
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func PullUser(c *gin.Context, email string) {
	validEmail := util.IsValidEmail(email)
	if validEmail {
		c.Set("Error", "E-mail is invalid")
		c.Status(http.StatusBadRequest)
		return
	}
	res, err := storage.ConsultUser(db.Repo, email)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func CreateUser(c *gin.Context, user inter.UserController) {
	valid := util.UltimateValidator(user.Email, user.Password)
	if valid == "" {
		wallet, err := lumx.PostLumxWallet()
		if err != nil {
			c.Set("Error", "Error in lumx Api")
			c.Status(http.StatusInternalServerError)
			return
		}
		newUser := inter.UserControllerInputDb{
			Email:     user.Email,
			Password:  user.Password,
			Wallet:    wallet.Address,
			WalletId:  wallet.ID,
			ProjectId: wallet.ProjectID,
		}
		res, err := storage.CreateUser(db.Repo, newUser)
		if err != nil {
			c.Set("Error", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Set("Response", res)
		c.Status(http.StatusOK)
	} else {
		c.Set("Error", valid)
		c.Status(http.StatusNotAcceptable)
		return
	}
}

func CreateContract(c *gin.Context, input inter.ContractController) {
	res, err := lumx.CreateContractLumx(input)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func DeployContract(c *gin.Context, input inter.DeployController) {
	res, err := lumx.DeployContract(input.ContractId)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}
