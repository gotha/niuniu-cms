package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	//dsn := "cms:4rfvbgt5@tcp(127.0.0.1:3306)/gocms?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := "host=localhost user=gocms password=4rfvbgt5 dbname=gocms port=5432 sslmode=disable TimeZone=Europe/Sofia"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Response{}, err
	}

	q := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	db.Exec(q)

	err = db.AutoMigrate(&Document{})
	if err != nil {
		return Response{}, err
	}

	doc := &Document{
		Title: "This is test",
		Body:  "This is my body",
	}

	res := db.Create(doc)
	fmt.Printf("%+v\n", res)

	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	resp, err := Handler(context.TODO())
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", err)
	//lambda.Start(Handler)
}
