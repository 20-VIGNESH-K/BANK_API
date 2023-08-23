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

type CustomerTableService struct {
	CustomerTableCollection *mongo.Collection
	ctx                     context.Context
}

func NewCustomerTableServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomerTable {
	return &CustomerTableService{collection, ctx}
}

func (t *CustomerTableService) CreateCustomerTable(user *models.CustomerTable) (string, error) {

	_, err := t.CustomerTableCollection.InsertOne(t.ctx, &user)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}

func (t *CustomerTableService) GetCustomerTable() ([]*models.CustomerTable, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.CustomerTableCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var customers []*models.CustomerTable
		for result.Next(ctx) {
			product := &models.CustomerTable{}
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
			return []*models.CustomerTable{}, nil
		}

		return customers, nil
	}
}

func (c *CustomerTableService) UpdateCustomerTable(id int64, customer *models.CustomerTable) (*mongo.UpdateResult, error) {
	iv := bson.M{"cus_id": id}
	fv := bson.M{"$set": &customer}
	res, err := c.CustomerTableCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CustomerTableService) DeleteCustomerTable(id int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "cus_id", Value: id}}

	res, err := c.CustomerTableCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}
