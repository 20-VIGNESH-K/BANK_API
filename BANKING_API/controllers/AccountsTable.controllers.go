package controllers

import (
	"banking-API/interfaces"
	"banking-API/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccountsTableController struct {
	AccountsTableService interfaces.IAccountsTable
}

func InitAccountsTableController(accountsTableService interfaces.IAccountsTable) AccountsTableController {
	return AccountsTableController{accountsTableService}
}

func (at *AccountsTableController) CreateAccountsTable(ctx *gin.Context) {
	var accountsTable *models.AccountsTable
	if err := ctx.ShouldBindJSON(&accountsTable); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newAccountsTable, err := at.AccountsTableService.CreateAccountsTable(accountsTable)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newAccountsTable})
}

func (pc *AccountsTableController) GetAccountsTable(ctx *gin.Context) {

	customers, err := pc.AccountsTableService.GetAccountsTable()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (c *AccountsTableController) UpdateAccountsTable(ctx *gin.Context) {
	id := ctx.Param("id")

	customer := &models.AccountsTable{}
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64) //str tp int
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := c.AccountsTableService.UpdateAccountsTable(id1, customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})

}

func (c *AccountsTableController) DeleteAccountsTable(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := c.AccountsTableService.DeleteAccountsTable(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
