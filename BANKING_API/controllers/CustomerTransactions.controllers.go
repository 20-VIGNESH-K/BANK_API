package controllers

import (
	"banking-API/interfaces"
	"banking-API/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerTransactionsController struct {
	CustomerTransactionsService interfaces.ICustomerTransactions
}

func InitCustomerTransactionsController(customerTransactionsService interfaces.ICustomerTransactions) CustomerTransactionsController {
	return CustomerTransactionsController{customerTransactionsService}
}

func (at *CustomerTransactionsController) CreateCustomerTransactions(ctx *gin.Context) {
	var customerTransactions *models.CustomerTransactions
	if err := ctx.ShouldBindJSON(&customerTransactions); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newCustomerTransactions, err := at.CustomerTransactionsService.CreateCustomerTransactions(customerTransactions)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newCustomerTransactions})
}

func (pc *CustomerTransactionsController) GetCustomerTransactions(ctx *gin.Context) {

	customers, err := pc.CustomerTransactionsService.GetCustomerTransactions()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (c *CustomerTransactionsController) UpdateCustomerTransactions(ctx *gin.Context) {
	id := ctx.Param("id")

	customer := &models.CustomerTransactions{}
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64) //str tp int
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := c.CustomerTransactionsService.UpdateCustomerTransactions(id1, customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})

}

func (c *CustomerTransactionsController) DeleteCustomerTransactions(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := c.CustomerTransactionsService.DeleteCustomerTransactions(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
