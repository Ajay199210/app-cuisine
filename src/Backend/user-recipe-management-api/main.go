package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"user-recipe-management.com/rest-api/db"
	"user-recipe-management.com/rest-api/routes"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	db.InitDB()
	server := gin.Default()

	//Enable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	server.Use(cors.New(config))

	server.POST("/userRecipes", routes.AddUserRecipe)
	server.GET("/userRecipes/:recipeId", routes.GetSingleUserRecipe)
	server.PUT("/userRecipes/:recipeId", routes.UpdateUserRecipe)
	server.DELETE("/userRecipes/:recipeId", routes.DeleteUserRecipe)

	server.GET("/allUserRecipes/:userId", routes.GetAllUserRecipesByUserID)
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
