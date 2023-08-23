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

type BankTableService struct {
	BankTableCollection *mongo.Collection
	ctx                     context.Context
}


func NewBankTableServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IBankTable {
	return &BankTableService{collection, ctx}
}


func (t *BankTableService) CreateBankTable(user *models.BankTable) (string, error) {

	_, err := t.BankTableCollection.InsertOne(t.ctx, &user)
	if err != nil {
		return "error", nil
	}

	return "success", nil
} 


func (t *BankTableService) GetBankTable() ([]*models.BankTable, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.BankTableCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var customers []*models.BankTable
		for result.Next(ctx) {
			product := &models.BankTable{}
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
			return []*models.BankTable{}, nil
		}

		return customers, nil
	}
}

func (c *BankTableService) UpdateBankTable(id int64, customer *models.BankTable) (*mongo.UpdateResult, error) {
	iv := bson.M{"bank_id": id}
	fv := bson.M{"$set": &customer}
	res, err := c.BankTableCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}





func (c *BankTableService) DeleteBankTable(id int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "bank_id", Value: id}}
	//var customer *models.Customer
	res, err := c.BankTableCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}




