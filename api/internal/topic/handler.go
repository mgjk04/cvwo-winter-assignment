package topic

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler interface {
	GetTopic(ctx *gin.Context)
	GetTopics(ctx *gin.Context)
	CreateTopic(ctx *gin.Context)
	UpdateTopic(ctx *gin.Context)
	DeleteTopic(ctx *gin.Context)
}

type topicHandler struct {
	s Service
}
//gotta rmb that methods implemented on type is just a function receiving that type,
//so gotta receive a pointer to handler cause its HUGE


//TODO: add logging 
func (h *topicHandler) GetTopic(ctx *gin.Context) {
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic ID"})
		return
	} 
	topic, err := h.s.FindByID(ctx, topicID)
	if err != nil {
		slog.Error(err.Error())
		switch err {
			case ErrTopicNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": "topic not found"})
				return
			case ErrTopicExists:
				ctx.JSON(http.StatusConflict, gin.H{"error": "topic already exists"})
				return
			case ErrUncaught:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
		}
	}
	
	ctx.JSON(http.StatusOK, topic)
}

func (h *topicHandler) GetTopics(ctx *gin.Context) {
	var query SearchQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameters"})
		return
	}
	topics, err := h.s.FindTopics(ctx, query.Page, query.Limit)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get topics"})
		return
	}
	ctx.JSON(http.StatusOK, topics)
}

func (h *topicHandler) CreateTopic(ctx *gin.Context){
	req := &TopicCreateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	topicID, err := h.s.CreateTopic(ctx, &Topic{TopicName: req.TopicName, Description: req.Description, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": topicID})
}
func (h *topicHandler) UpdateTopic(ctx *gin.Context){
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic ID"})
		return
	}
	req := &TopicUpdateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err = h.s.UpdateTopic(ctx, &Topic{ID: topicID, TopicName: req.TopicName, Description: req.Description, AuthorID: req.AuthorID})
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update topic"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *topicHandler) DeleteTopic(ctx *gin.Context){
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic ID"})
		return
	}
	err = h.s.DeleteTopic(ctx, topicID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete topic"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewTopicHandler(s Service) *topicHandler {
	return &topicHandler{s: s}
}