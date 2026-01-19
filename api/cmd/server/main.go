package main

import (
	"context"
	"log/slog"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/auth"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/comment"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/middleware"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/post"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/topic"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/user"
	"github.com/mgjk04/cvwo-winter-assignment/api/pkg/db"
)

func main() {
	//env variables
	addr := os.Getenv("ADDR")
	dbURL := os.Getenv("DB_URL")
	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")
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
	authSvc := auth.NewAuthSvc(userSvc, accessSecret, refreshSecret)
	//init handlers
	userHandler := user.NewUserHandler(userSvc)
	topicHandler := topic.NewTopicHandler(topicSvc)
	postHandler := post.NewPostHandler(postSvc)
	commentHandler := comment.NewCommentHandler(commentSvc)
	authHandler := auth.NewAuthHandler(authSvc)

	//auth middleware
	authMiddleware := middleware.AuthMiddleware(authSvc)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"}, //remember to change the domain
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
    AllowCredentials: true,}))
	r.Use(middleware.ErrorHandler())

	//auth
	r.POST("/login",authHandler.Login)
	r.POST("/logout",authHandler.Logout)
	r.POST("/signup",authHandler.Signup)

	r.POST("/refresh", authHandler.Refresh)

	//users
	users := r.Group("/users")
	users.GET("/:userId", userHandler.GetUser)
	users.POST("/", authMiddleware, userHandler.CreateUser)
	users.DELETE("/:userId", authMiddleware, userHandler.DeleteUser)

	//topics
	topics := r.Group("/topics")
	topics.GET("/:topicId", topicHandler.GetTopic)
	topics.GET("/", topicHandler.GetTopics)
	topics.POST("/", authMiddleware, topicHandler.CreateTopic)
	topics.PUT("/:topicId",authMiddleware, topicHandler.UpdateTopic)
	topics.DELETE("/:topicId", authMiddleware, topicHandler.DeleteTopic)

	//posts
	topics.GET("/:topicId/posts", postHandler.GetPosts)
	topics.POST("/:topicId/posts", postHandler.CreatePost)

	posts := r.Group("/posts")
	posts.GET("/:postId", postHandler.GetPost)
	posts.PUT("/:postId", authMiddleware, postHandler.UpdatePost)
	posts.DELETE("/:postId",authMiddleware, postHandler.DeletePost)

	//comments
	posts.GET("/:postId/comments", commentHandler.GetComments)
	posts.POST("/:postId/comments", authMiddleware, commentHandler.CreateComment)

	comments := r.Group("/comments")
	comments.GET("/:commentId", commentHandler.GetComment)
	comments.PUT("/:commentId", authMiddleware, commentHandler.UpdateComment)
	comments.DELETE("/:commentId", authMiddleware, commentHandler.DeleteComment)
	r.Run(addr)
}