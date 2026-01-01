package comment

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
		return
	} 
	comment, err := h.s.FindByID(ctx, commentID)
	if err != nil {
		slog.Error(err.Error())
		switch err {
			case ErrCommentNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
				return
			case ErrUncaught:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
		}
	}
	
	ctx.JSON(http.StatusOK, comment)
}

func (h *commentHandler) GetComments(ctx *gin.Context) {
	var query SearchQuery
	postID, err := uuid.Parse(ctx.Param("postId"))
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}
	comments, err := h.s.FindComments(ctx, postID, query.Page, query.Limit)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get comments"})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (h *commentHandler) CreateComment(ctx *gin.Context){
	postID, err := uuid.Parse(ctx.Param("postId"))
	req := &CommentCreateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	commentID, err := h.s.CreateComment(ctx, &Comment{Content: req.Content, PostID: postID, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": commentID})
}
func (h *commentHandler) UpdateComment(ctx *gin.Context){
	commentID, err := uuid.Parse(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
		return
	}
	req := &CommentUpdateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err = h.s.UpdateComment(ctx, &Comment{ID: commentID, Content: req.Content, PostID: req.PostID, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update comment"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *commentHandler) DeleteComment(ctx *gin.Context){
	commentID, err := uuid.Parse(ctx.Param("commentId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}
	err = h.s.DeleteComment(ctx, commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete comment"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewCommentHandler(s Service) *commentHandler {
	return &commentHandler{s: s}
}