package routes

import (
	"banking-API/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(router *gin.Engine) {
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is Healthy"})
	})
}

func CustomerTableRoute(router *gin.Engine, controller controllers.CustomerTableController) {

	router.POST("/api/bankAccount/createCustomerTable", controller.CreateCustomerTable)
	router.PUT("/api/bankAccount/updateCustomerTable/:id", controller.UpdateCustomerTable)
	router.GET("/api/bankAccount/getCustomerTable", controller.GetCustomerTable)
	router.DELETE("/api/bankAccount/deleteCustomerTable/:id", controller.DeleteCustomerTable)
}

func CustomerTransactionsRoute(router *gin.Engine, controller controllers.CustomerTransactionsController) {

	router.POST("/api/bankAccount/createCustomerTransactions", controller.CreateCustomerTransactions)
	router.PUT("/api/bankAccount/updateCustomerTransactions/:id", controller.UpdateCustomerTransactions)
	router.GET("/api/bankAccount/getCustomerTransactions", controller.GetCustomerTransactions)
	router.DELETE("/api/bankAccount/deleteCustomerTransactions/:id", controller.DeleteCustomerTransactions)
}

func LoanCollectionsRoute(router *gin.Engine, controller controllers.LoanCollectionsController) {

	router.POST("/api/bankAccount/createLoanCollections", controller.CreateLoanCollections)
	router.PUT("/api/bankAccount/updateLoanCollections/:id", controller.UpdateLoanCollections)
	router.GET("/api/bankAccount/getLoanCollections", controller.GetLoanCollections)
	router.DELETE("/api/bankAccount/deleteLoanCollections/:id", controller.DeleteLoanCollections)
}

func AccountsTableRoute(router *gin.Engine, controller controllers.AccountsTableController) {

	router.POST("/api/bankAccount/createAccount", controller.CreateAccountsTable)
	router.PUT("/api/bankAccount/updateAccount/:id", controller.UpdateAccountsTable)
	router.GET("/api/bankAccount/getAccount", controller.GetAccountsTable)
	router.DELETE("/api/bankAccount/deleteAccount/:id", controller.DeleteAccountsTable)
}

func BankTableRoute(router *gin.Engine, controller controllers.BankTableController) {

	router.POST("/api/bankAccount/createBankTable", controller.CreateBankTable)
	router.PUT("/api/bankAccount/updateBankTable/:id", controller.UpdateBankTable)
	router.GET("/api/bankAccount/getBankTable", controller.GetBankTable)
	router.DELETE("/api/bankAccount/deleteBankTable/:id", controller.DeleteBankTable)

}
