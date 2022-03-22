package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Level store level model
type Level struct {
	Id     string  `json:"-"`
	Levels [][]int `json:"levels"`
}

//GetById retrieve level by id from AWS
func GetById(db *dynamodb.DynamoDB, id string) (*Level, error) {
	var l Level

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("greenjade_task"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &l)

	if err != nil {
		return nil, err
	}

	return &l, nil
}

//Store store new item
func Store(db *dynamodb.DynamoDB, id int, levels [][]int) (int, error) {

	item := struct {
		Id     *int    `json:"id"`
		Levels [][]int `json:"levels"`
	}{
		Id:     aws.Int(id),
		Levels: levels,
	}

	av, _ := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("greenjade_task"),
	}

	_, err := db.PutItem(input)

	if err != nil {
		return 0, err
	}

	return id, nil
}

//Update update level by id
func Update(db *dynamodb.DynamoDB, id string, levels [][]int) error {

	item := struct {
		Levels [][]int `json:"levels"`
	}{
		Levels: levels,
	}

	av, _ := dynamodbattribute.MarshalMap(item)

	levelInput := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("greenjade_task"),
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":levels": levelInput.Item["levels"],
		},
		TableName: aws.String("greenjade_task"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		UpdateExpression: aws.String("set levels = :levels"),
	}

	_, err := db.UpdateItem(input)

	if err != nil {
		return err
	}

	return nil
}

//Remove remove item
func Remove(db *dynamodb.DynamoDB, id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("greenjade_task"),
	}

	_, err := db.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
