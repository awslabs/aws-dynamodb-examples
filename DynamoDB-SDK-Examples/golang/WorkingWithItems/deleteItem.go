package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    
    "fmt"
    "os"
)

// Item struct to hold keys
type Item struct {
    Pk string  `json:"pk"`
    Sk string  `json:"sk"`
}

func main() {
  
	// Create Session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("eu-west-1")},
    )

    // Create DynamoDB client
    svc := dynamodb.New(sess)

	// Keys for item
    item := Item{
        Pk: "jose.schneller@somewhere.com",
        Sk: "metadata",
    }

	// Marshal
    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        fmt.Println("Got error marshalling map:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

	// Delete Item
    input := &dynamodb.DeleteItemInput{
		TableName: aws.String("RetailDatabase"),
        Key:       av,
    }

    _, err = svc.DeleteItem(input)
    if err != nil {
        fmt.Println("Got error calling DeleteItem")
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Deleted Item")
}