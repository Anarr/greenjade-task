package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//Level store level model
type Level struct {
	Id string `json:"id"`
	Levels [][]int `json:"levels"`
}

//WithEmptyId chck  returned data id is empty or not
func (l *Level) WithEmptyId() bool {
	return l.Id == ""
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