package controllers

import (
	"banking-API/interfaces"
	"banking-API/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoanCollectionsController struct {
	LoanCollectionsService interfaces.ILoanCollections
}

func InitLoanCollectionsController(loanCollectionsService interfaces.ILoanCollections) LoanCollectionsController {
	return LoanCollectionsController{loanCollectionsService}
}

func (at *LoanCollectionsController) CreateLoanCollections(ctx *gin.Context) {
	var loanCollections *models.LoanCollections
	if err := ctx.ShouldBindJSON(&loanCollections); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newLoanCollections, err := at.LoanCollectionsService.CreateLoanCollections(loanCollections)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newLoanCollections})
}

func (pc *LoanCollectionsController) GetLoanCollections(ctx *gin.Context) {

	customers, err := pc.LoanCollectionsService.GetLoanCollections()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (c *LoanCollectionsController) UpdateLoanCollections(ctx *gin.Context) {
	id := ctx.Param("id")

	customer := &models.LoanCollections{}
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id1, err := strconv.ParseInt(id, 10, 64) //str tp int
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := c.LoanCollectionsService.UpdateLoanCollections(id1, customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})

}

func (c *LoanCollectionsController) DeleteLoanCollections(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := c.LoanCollectionsService.DeleteLoanCollections(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}
