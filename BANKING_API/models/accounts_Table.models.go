package models

type CustomerTable struct {
	ID         int64  `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	Password   string `json:"password" bson:"password"`
	CustomerID int64  `json:"cus_id" bson:"cus_id"`
	BankID     int64  `json:"bank_id" bson:"bank_id"`
}

type CustomerTransactions struct {
	CustomerID        int64   `json:"cus_id" bson:"cus_id"`
	TransactionAmount float64 `json:"transaction_amount" bson:"transaction_amount"`
}

type LoanCollections struct {
	CustomerID int64   `json:"cus_id" bson:"cus_id"`
	LoanAmount float64 `json:"loan_amount" bson:"loan_amount"`
	LoanType   string  `json:"loan_type" bson:"loan_type"`
}

type AccountsTable struct {
	AccountID   int64  `json:"acc_id" bson:"_acc_id"`
	CustomerID  int64  `json:"cus_id" bson:"cus_id"`
	AccountType string `json:"acc_type" bson:"acc_type"`
	BranchAddr  string `json:"acc_branch_address" bson:"acc_branch_address"`
}

type BankTable struct {
	BankID      int64  `json:"bank_id" bson:"bank_id"`
	IFSCCode    string `json:"ifsc_code" bson:"ifsc_code"`
	BankAddress string `json:"bank_address" bson:"bank_address"`
}
