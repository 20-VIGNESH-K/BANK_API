package controllers

import (
	"banking-API/interfaces"
	"banking-API/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerTableController struct {
	CustomerTableService interfaces.ICustomerTable
}

func InitCustomerTableController(customerTableService interfaces.ICustomerTable) CustomerTableController {
	return CustomerTableController{customerTableService}
}

func (at *CustomerTableController) CreateCustomerTable(ctx *gin.Context) {
	var customerTable *models.CustomerTable
	if err := ctx.ShouldBindJSON(&customerTable); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newCustomerTable, err := at.CustomerTableService.CreateCustomerTable(customerTable)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newCustomerTable})
}

func (pc *CustomerTableController) GetCustomerTable(ctx *gin.Context) {

	customers, err := pc.CustomerTableService.GetCustomerTable()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (c *CustomerTableController) UpdateCustomerTable(ctx *gin.Context) {
	id := ctx.Param("id")

	customer := &models.CustomerTable{}
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64) //str tp int
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := c.CustomerTableService.UpdateCustomerTable(id1, customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})

}

func (c *CustomerTableController) DeleteCustomerTable(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := c.CustomerTableService.DeleteCustomerTable(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
