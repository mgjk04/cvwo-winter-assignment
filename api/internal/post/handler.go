package post

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Handler interface {
	GetPost(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
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
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	post, err := h.s.FindByID(ctx, postID)
	if err != nil {
		ctx.Error(err)
	}
	
	ctx.JSON(http.StatusOK, post)
}

func (h *postHandler) GetPosts(ctx *gin.Context) {
	var query SearchQuery
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	posts, err := h.s.FindPosts(ctx, topicID, query.Page, query.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, &PostReadRes{Posts: posts, Count: len(posts)})
}

func (h *postHandler) CreatePost(ctx *gin.Context){
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.Error(generalErrors.ErrUnauthorized)
		return
	}
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	req := &PostCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	postID, err := h.s.CreatePost(ctx, &Post{Title: req.Title, Description: req.Description, TopicID: topicID, AuthorID: parsedUserID})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": postID})
}
func (h *postHandler) UpdatePost(ctx *gin.Context){
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.Error(generalErrors.ErrUnauthorized)
		return
	}
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	req := &PostUpdateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	var finalAuthorID uuid.UUID
	if req.AuthorID != uuid.Nil {
		finalAuthorID = req.AuthorID
	} else {
		finalAuthorID = parsedUserID
	}
	err = h.s.UpdatePost(ctx, parsedUserID, &Post{ID: postID, Title: req.Title, Description: req.Description, TopicID: req.TopicID, AuthorID: finalAuthorID})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *postHandler) DeletePost(ctx *gin.Context){
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.Error(generalErrors.ErrUnauthorized)
		return
	}
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	err = h.s.DeletePost(ctx, parsedUserID, postID)
	if err != nil {
		ctx.Error(err)	
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewPostHandler(s Service) *postHandler {
	return &postHandler{s: s}
}