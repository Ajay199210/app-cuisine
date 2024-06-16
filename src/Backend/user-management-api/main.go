package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"usermanagement.com/rest-api/db"
	"usermanagement.com/rest-api/routes"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	db.InitDB()
	server := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	server.Use(cors.New(config))

	server.POST("/signup", routes.Signup)
	server.POST("/login", routes.Login)
	server.GET("/logout", routes.Logout)
	server.GET("/users", routes.GetUsers)
	server.GET("/users/:id", routes.GetUser)
	server.PUT("/users/:id", routes.UpdateUser)
	server.DELETE("/users/:id", routes.DeleteUser)

	ginLambda = ginadapter.NewV2(server)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// Log the request path
	fmt.Printf("Request path: %s\n", req.RawPath)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
