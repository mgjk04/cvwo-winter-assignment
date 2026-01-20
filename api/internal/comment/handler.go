package comment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Handler interface {
	GetComment(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandler struct {
	s Service
}
//gotta rmb that methods implemented on type is just a function receiving that type,
//so gotta receive a pointer to handler cause its HUGE


//TODO: add logging 
func (h *commentHandler) GetComment(ctx *gin.Context) {
	commentID, err := uuid.Parse(ctx.Param("commentId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	comment, err := h.s.FindByID(ctx, commentID)
	if err != nil {
		ctx.Error(err)
		return
	}
	
	ctx.JSON(http.StatusOK, comment)
}

func (h *commentHandler) GetComments(ctx *gin.Context) {
	var query SearchQuery
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	comments, err := h.s.FindComments(ctx, postID, query.Page, query.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, &CommentReadRes{Comments: comments, Count: len(comments)})
}

func (h *commentHandler) CreateComment(ctx *gin.Context){
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
	req := &CommentCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	commentID, err := h.s.CreateComment(ctx, &Comment{Content: req.Content, PostID: postID, AuthorID: parsedUserID})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": commentID})
}
func (h *commentHandler) UpdateComment(ctx *gin.Context){
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
	commentID, err := uuid.Parse(ctx.Param("commentId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	req := &CommentUpdateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	err = h.s.UpdateComment(ctx, parsedUserID, &Comment{ID: commentID, Content: req.Content, PostID: req.PostID, AuthorID: parsedUserID})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *commentHandler) DeleteComment(ctx *gin.Context){
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
	commentID, err := uuid.Parse(ctx.Param("commentId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	err = h.s.DeleteComment(ctx, parsedUserID, commentID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewCommentHandler(s Service) *commentHandler {
	return &commentHandler{s: s}
}