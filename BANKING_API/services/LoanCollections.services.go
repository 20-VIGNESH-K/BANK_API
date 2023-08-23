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

type LoanCollectionsService struct {
	LoanCollectionsCollection *mongo.Collection
	ctx                     context.Context
}

func NewLoanCollectionsServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ILoanCollections {
	return &LoanCollectionsService{collection, ctx}
}

func (t *LoanCollectionsService) CreateLoanCollections(user *models.LoanCollections) (string, error) {

	_, err := t.LoanCollectionsCollection.InsertOne(t.ctx, &user)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}

func (t *LoanCollectionsService) GetLoanCollections() ([]*models.LoanCollections, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{}}
	result, err := t.LoanCollectionsCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var customers []*models.LoanCollections
		for result.Next(ctx) {
			product := &models.LoanCollections{}
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
			return []*models.LoanCollections{}, nil
		}

		return customers, nil
	}
}

func (c *LoanCollectionsService) UpdateLoanCollections(id int64, customer *models.LoanCollections) (*mongo.UpdateResult, error) {
	iv := bson.M{"cus_id": id}
	fv := bson.M{"$set": &customer}
	res, err := c.LoanCollectionsCollection.UpdateOne(c.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *LoanCollectionsService) DeleteLoanCollections(id int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{Key: "cus_id", Value: id}}

	res, err := c.LoanCollectionsCollection.DeleteOne(c.ctx, filter)

	if err != nil {
		return nil, err
	}
	return res, nil
}
