package main

import (
	"banking-API/config"
	"banking-API/constants"
	"banking-API/controllers"
	"banking-API/routes"
	"banking-API/services"

	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()

	accountsTableCollection := mongoClient.Database(constants.DatabaseName).Collection("AccountsTable")
	accountsTableService := services.NewAccountsTableServiceInit(accountsTableCollection, ctx)
	accountsTableController := controllers.InitAccountsTableController(accountsTableService)
	routes.AccountsTableRoute(server, accountsTableController)

	customerTableCollection := mongoClient.Database(constants.DatabaseName).Collection("CustomerTable")
	customerTableService := services.NewCustomerTableServiceInit(customerTableCollection, ctx)
	customerTableController := controllers.InitCustomerTableController(customerTableService)
	routes.CustomerTableRoute(server, customerTableController)

	customerTransactionsCollection := mongoClient.Database(constants.DatabaseName).Collection("CustomerTransactions")
	customerTransactionsService := services.NewCustomerTransactionsServiceInit(customerTransactionsCollection, ctx)
	customerTransactionsController := controllers.InitCustomerTransactionsController(customerTransactionsService)
	routes.CustomerTransactionsRoute(server, customerTransactionsController)

	loanCollectionsCollection := mongoClient.Database(constants.DatabaseName).Collection("LoanCollections")
	loanCollectionsService := services.NewLoanCollectionsServiceInit(loanCollectionsCollection, ctx)
	loanCollectionsController := controllers.InitLoanCollectionsController(loanCollectionsService)
	routes.LoanCollectionsRoute(server, loanCollectionsController)

	bankTableCollection := mongoClient.Database(constants.DatabaseName).Collection("BankTable")
	bankTableService := services.NewBankTableServiceInit(bankTableCollection, ctx)
	bankTableController := controllers.InitBankTableController(bankTableService)
	routes.BankTableRoute(server, bankTableController)

}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))

}
