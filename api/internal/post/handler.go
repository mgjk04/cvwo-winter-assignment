package post

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler interface {
	GetPost(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

type postHandler struct {
	s Service
}
//gotta rmb that methods implemented on type is just a function receiving that type,
//so gotta receive a pointer to handler cause its HUGE


//TODO: add logging 
func (h *postHandler) GetPost(ctx *gin.Context) {
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	} 
	post, err := h.s.FindByID(ctx, postID)
	if err != nil {
		slog.Error(err.Error())
		switch err {
			case ErrPostNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
				return
			case ErrUncaught:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
		}
	}
	
	ctx.JSON(http.StatusOK, post)
}

func (h *postHandler) GetPosts(ctx *gin.Context) {
	var query SearchQuery
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}
	posts, err := h.s.FindPosts(ctx, topicID, query.Page, query.Limit)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get posts"})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (h *postHandler) CreatePost(ctx *gin.Context){
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	req := &PostCreateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	postID, err := h.s.CreatePost(ctx, &Post{Title: req.Title, Description: req.Description, TopicID: topicID, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": postID})
}
func (h *postHandler) UpdatePost(ctx *gin.Context){
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	req := &PostUpdateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err = h.s.UpdatePost(ctx, &Post{ID: postID, Title: req.TopicName, Description: req.Description, TopicID: req.TopicID, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update topic"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "post updated successfully"})
}
func (h *postHandler) DeletePost(ctx *gin.Context){
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	err = h.s.DeletePost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
}

func NewPostHandler(s Service) *postHandler {
	return &postHandler{s: s}
}