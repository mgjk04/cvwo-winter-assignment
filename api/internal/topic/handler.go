package topic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
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
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	topic, err := h.s.FindByID(ctx, topicID)
	if err != nil {
		ctx.Error(err)
		return
	}
	
	ctx.JSON(http.StatusOK, topic)
}

func (h *topicHandler) GetTopics(ctx *gin.Context) {
	var query SearchQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	topics, err := h.s.FindTopics(ctx, query.Page, query.Limit)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, topics)
}

func (h *topicHandler) CreateTopic(ctx *gin.Context){
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
	req := &TopicCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	topicID, err := h.s.CreateTopic(ctx, &Topic{TopicName: req.TopicName, Description: req.Description, AuthorID: parsedUserID})
	if err != nil {
		ctx.Error(err);
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": topicID})
}
func (h *topicHandler) UpdateTopic(ctx *gin.Context){
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.Error(generalErrors.ErrUnauthorized)
	}
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}

	//verify ownership
	req := &TopicUpdateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	err = h.s.UpdateTopic(ctx, parsedUserID, &Topic{ID: topicID, TopicName: req.TopicName, Description: req.Description, AuthorID: req.AuthorID})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
func (h *topicHandler) DeleteTopic(ctx *gin.Context){
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.Error(generalErrors.ErrUnauthorized)
	}
	parsedUserID, err := uuid.Parse(userID.(string))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	topicID, err := uuid.Parse(ctx.Param("topicId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	
	err = h.s.DeleteTopic(ctx, parsedUserID, topicID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewTopicHandler(s Service) *topicHandler {
	return &topicHandler{s: s}
}