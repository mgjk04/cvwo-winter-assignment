package main

import (
	"context"
	"log/slog"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/comment"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/post"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/topic"
	"github.com/mgjk88/cvwo-winter-assignment/api/internal/user"
	"github.com/mgjk88/cvwo-winter-assignment/api/pkg/db"
)

func main() {
	//env variables
	addr := os.Getenv("ADDR")
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
	postRepo := post.NewPostRepo(pool)
	commentRepo := comment.NewCommentRepo(pool)
	//init svcs
	userSvc := user.NewUserSvc(userRepo)
	topicSvc := topic.NewTopicSvc(topicRepo)
	postSvc := post.NewPostSvc(postRepo)
	commentSvc := comment.NewCommentSvc(commentRepo)
	//init handlers
	userHandler := user.NewUserHandler(userSvc)
	topicHandler := topic.NewTopicHandler(topicSvc)
	postHandler := post.NewPostHandler(postSvc)
	commentHandler := comment.NewCommentHandler(commentSvc)


	r := gin.Default()
	//users
	users := r.Group("/users")
	users.GET("/:userId", userHandler.GetUser)
	users.POST("/", userHandler.CreateUser)
	users.DELETE("/:userId", userHandler.DeleteUser)

	//topics
	topics := r.Group("/topics")
	topics.GET("/:topicId", topicHandler.GetTopic)
	topics.GET("/", topicHandler.GetTopics)
	topics.POST("/", topicHandler.CreateTopic)
	topics.PUT("/:topicId", topicHandler.UpdateTopic)
	topics.DELETE("/:topicId", topicHandler.DeleteTopic)

	//posts
	topics.GET("/:topicId/posts", postHandler.GetPosts)
	topics.POST("/:topicId/posts", postHandler.CreatePost)

	posts := r.Group("/posts")
	posts.GET("/:postId", postHandler.GetPost)
	posts.PUT("/:postId", postHandler.UpdatePost)
	posts.DELETE("/:postId", postHandler.DeletePost)

	//comments
	posts.GET("/:postId/comments", commentHandler.GetComments)
	posts.POST("/:postId/comments", commentHandler.CreateComment)

	comments := r.Group("/comments")
	comments.GET("/:commentId", commentHandler.GetComment)
	comments.PUT("/:commentId", commentHandler.UpdateComment)
	comments.DELETE("/:commentId", commentHandler.DeleteComment)
	r.Run(addr)
}