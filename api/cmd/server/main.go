package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/topic"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/user"
	"github.com/mgjk88/cvwo-winter-assignment/api/pkg/db"
)

//TODO: implement logging using slog later
func main() {
	//env variables
	//addr := os.Getenv("ADDR")
	dbURL := os.Getenv("DB_URL")
	// env := os.Getenv("ENV")
	// if env == "DEV" {
	// 	gin.SetMode(gin.DebugMode)
	// }

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        AddSource: true, // show file and line number
        Level:     slog.LevelDebug,
    }))
	slog.SetDefault(logger)

	pool, err := db.NewPool(context.Background(), dbURL)
	if err != nil {
		//well we're cooked
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer pool.Close()

	//init repos
	userRepo := user.NewUserRepo(pool)
	topicRepo := topic.NewTopicRepo(pool)
	//init svcs
	userSvc := user.NewUserSvc(userRepo)
	topicSvc := topic.NewTopicSvc(topicRepo)
	//init handlers
	userHandler := user.NewUserHandler(userSvc)
	topicHandler := topic.NewTopicHandler(topicSvc)


	r := gin.Default()
	users := r.Group("/users")
	users.GET("/:userId", userHandler.GetUser)
	users.POST("/", userHandler.CreateUser)
	users.DELETE("/:userId", userHandler.DeleteUser)

	topics := r.Group("/topics")
	topics.GET("/:topicId", topicHandler.GetTopic)
	topics.GET("/", topicHandler.GetTopics)
	topics.POST("/", topicHandler.CreateTopic)
	topics.PUT("/:topicId", topicHandler.UpdateTopic)
	topics.DELETE("/:topicId", topicHandler.DeleteTopic)
	r.Run()
}