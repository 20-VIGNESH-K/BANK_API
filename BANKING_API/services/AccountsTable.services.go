package services

import (
	"banking-API/interfaces"
	"banking-API/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountsTableService struct {
	AccountsTableCollection *mongo.Collection
	ctx                     context.Context
}


func NewAccountsTableServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IAccountsTable {
	return &AccountsTableService{collection, ctx}
}


func (t *AccountsTableService) CreateAccountsTable(user *models.AccountsTable) (string, error) {

	_, err := t.AccountsTableCollection.InsertOne(t.ctx, &user)
	if err != nil {
		return "error", nil
	}

	return "success", nil
} 


func (t *AccountsTableService) GetAccountsTable() ([]*models.AccountsTable, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.AccountsTableCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var customers []*models.AccountsTable
		for result.Next(ctx) {
			product := &models.AccountsTable{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}
			customers = append(customers, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(customers) == 0 {
			return []*models.AccountsTable{}, nil
		}

		return customers, nil
	}
}

func (c *AccountsTableService) UpdateAccountsTable(id int64, customer *models.AccountsTable) (*mongo.UpdateResult, error) {
	iv := bson.M{"cus_id": id}
	fv := bson.M{"$set": &customer}
	res, err := c.AccountsTableCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}





func (c *AccountsTableService) DeleteAccountsTable(id int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "cus_id", Value: id}}
	//var customer *models.Customer
	res, err := c.AccountsTableCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}




