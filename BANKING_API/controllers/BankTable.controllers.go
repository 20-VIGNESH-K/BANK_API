package controllers

import (
	"banking-API/interfaces"
	"banking-API/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BankTableController struct {
	BankTableService interfaces.IBankTable
}

func InitBankTableController(bankTableService interfaces.IBankTable) BankTableController {
	return BankTableController{bankTableService}
}

func (at *BankTableController) CreateBankTable(ctx *gin.Context) {
	var bankTable *models.BankTable
	if err := ctx.ShouldBindJSON(&bankTable); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newBankTable, err := at.BankTableService.CreateBankTable(bankTable)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBankTable})
}

func (pc *BankTableController) GetBankTable(ctx *gin.Context) {

	customers, err := pc.BankTableService.GetBankTable()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (c *BankTableController) UpdateBankTable(ctx *gin.Context) {
	id := ctx.Param("id")

	customer := &models.BankTable{}
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64) //str tp int
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := c.BankTableService.UpdateBankTable(id1, customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})

}


func (c * BankTableController) DeleteBankTable(ctx *gin.Context){
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := c.BankTableService.DeleteBankTable(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
} 
