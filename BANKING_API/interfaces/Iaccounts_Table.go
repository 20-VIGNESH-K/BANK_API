package interfaces

import (
	"banking-API/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICustomerTable interface {
	CreateCustomerTable(CustomerTable *models.CustomerTable) (string, error)
	UpdateCustomerTable(id int64, updatedCustomerTable *models.CustomerTable) (*mongo.UpdateResult, error)
	DeleteCustomerTable(id int64) (*mongo.DeleteResult, error)
	GetCustomerTable() ([]*models.CustomerTable, error)
}

type ICustomerTransactions interface {
	CreateCustomerTransactions(CustomerTransactions *models.CustomerTransactions) (string, error)
	UpdateCustomerTransactions(id int64, updatedCustomerTransactions *models.CustomerTransactions) (*mongo.UpdateResult, error)
	DeleteCustomerTransactions(id int64) (*mongo.DeleteResult, error)
	GetCustomerTransactions() ([]*models.CustomerTransactions, error)
}

type ILoanCollections interface {
	CreateLoanCollections(LoanCollections *models.LoanCollections) (string, error)
	UpdateLoanCollections(id int64, updatedLoanCollections *models.LoanCollections) (*mongo.UpdateResult, error)
	DeleteLoanCollections(id int64) (*mongo.DeleteResult, error)
	GetLoanCollections() ([]*models.LoanCollections, error)
}

type IAccountsTable interface {
	CreateAccountsTable(AccountsTable *models.AccountsTable) (string, error)
	UpdateAccountsTable(id int64, updatedAccountsTable *models.AccountsTable) (*mongo.UpdateResult, error)
	DeleteAccountsTable(id int64) (*mongo.DeleteResult, error)
	GetAccountsTable() ([]*models.AccountsTable, error)
}

type IBankTable interface {
	CreateBankTable(BankTable *models.BankTable) (string, error)
	UpdateBankTable(id int64, updatedBankTable *models.BankTable) (*mongo.UpdateResult, error)
	DeleteBankTable(id int64) (*mongo.DeleteResult, error)
	GetBankTable() ([]*models.BankTable, error)
}
