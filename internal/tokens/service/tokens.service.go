package service

import (
	inter "fanify/internal/tokens/interface"
	"fanify/package/lumx"
	lInterface "fanify/package/lumx/interface"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, input inter.TokenCreateInput) {
	newInput := lInterface.TokenTypeRequest{
		Traits:      *input.Traits,
		MaxSupply:   input.MaxSupply,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
	}
	res, err := lumx.CreateTokenType(input.ContractId, newInput)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func MakeTransfer(c *gin.Context, input inter.Transaction) {
	newInput := lInterface.TransferDetails{
		ContractId: input.ContractID,
		From:       input.From,
		To:         input.To,
		Quantity:   input.Quantity,
	}
	_, err := lumx.CreateTransfer(newInput)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", "Transfer is by made. It may take a few minutes if you have a lot of orders, but it will be delivered on time :)")
	c.Status(http.StatusOK)
}

func MakeTrade(c *gin.Context, input inter.Trade) {
	newInput := lInterface.TransferDetails{
		ContractId: input.Transaction1.ContractID,
		From:       input.Transaction1.From,
		To:         input.Transaction1.To,
		Quantity:   input.Transaction1.Quantity,
	}
	_, err := lumx.CreateTransfer(newInput)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	newInput = lInterface.TransferDetails{
		ContractId: input.Transaction2.ContractID,
		From:       input.Transaction2.From,
		To:         input.Transaction2.To,
		Quantity:   input.Transaction2.Quantity,
	}
	_, err = lumx.CreateTransfer(newInput)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", "Transfer is by made. It may take a few minutes if you have a lot of orders, but it will be delivered on time :)")
	c.Status(http.StatusOK)
}

func GainTokens(c *gin.Context, input inter.Gimmetoken) {
	randomNumber := rand.Intn(1000000) + 1
	newInput := lInterface.MintRequest{
		ContractID: input.ContractID,
		WalletID:   input.WalletID,
		Quantity:   randomNumber,
		URINumber:  input.URINumber,
	}
	res, err := lumx.MintTransaction(newInput)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func GetToken(c *gin.Context, UriToken string, ContractId string) {
	num, err := strconv.Atoi(UriToken)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	res, err := lumx.FetchTokenType(ContractId, num)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func GetAllToken(c *gin.Context, ContractId string) {
	res, err := lumx.FetchTokenTypes(ContractId)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}
