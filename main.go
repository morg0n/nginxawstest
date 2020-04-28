package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetAndUpdateCount() (uint64, error) {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// Create item in table Movies
	tableName := "Movies"
	movieName := "The Big New Movie"
	movieYear := "2015"
	movieRating := "0.5"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(movieRating),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(movieYear),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)
	return 1, nil

}

func HandleRequest(ctx context.Context) (string, error) {

	return fmt.Sprintln("I'm crazy Golang!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}