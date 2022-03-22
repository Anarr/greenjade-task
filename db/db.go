package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DatabaseConnInterface interface {
	Connect() (*dynamodb.DynamoDB, error)
}

type AwsConnection struct {
	ApiKey string
	SecretKey string
	Region string
}

func (ac AwsConnection) Connect() (*dynamodb.DynamoDB, error) {
	creds := credentials.NewStaticCredentials(ac.ApiKey, ac.SecretKey, "")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(ac.Region),
			Credentials: creds,
		},

	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return svc, nil
}

//NewConnection create new database connection
func NewConnection(database DatabaseConnInterface) (*dynamodb.DynamoDB, error) {
	return database.Connect()
}