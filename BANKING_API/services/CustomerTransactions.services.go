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

type CustomerTransactionsService struct {
	CustomerTransactionsCollection *mongo.Collection
	ctx                     context.Context
}

func NewCustomerTransactionsServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomerTransactions {
	return &CustomerTransactionsService{collection, ctx}
}

func (t *CustomerTransactionsService) CreateCustomerTransactions(user *models.CustomerTransactions) (string, error) {

	_, err := t.CustomerTransactionsCollection.InsertOne(t.ctx, &user)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}

func (t *CustomerTransactionsService) GetCustomerTransactions() ([]*models.CustomerTransactions, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.CustomerTransactionsCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var customers []*models.CustomerTransactions
		for result.Next(ctx) {
			product := &models.CustomerTransactions{}
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
			return []*models.CustomerTransactions{}, nil
		}

		return customers, nil
	}
}

func (c *CustomerTransactionsService) UpdateCustomerTransactions(id int64, customer *models.CustomerTransactions) (*mongo.UpdateResult, error) {
	iv := bson.M{"cus_id": id}
	fv := bson.M{"$set": &customer}
	res, err := c.CustomerTransactionsCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CustomerTransactionsService) DeleteCustomerTransactions(id int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "cus_id", Value: id}}

	res, err := c.CustomerTransactionsCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}
