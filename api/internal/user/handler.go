package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Handler interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	s Service
}
//gotta rmb that methods implemented on type is just a function receiving that type,
//so gotta receive a pointer to handler cause its HUGE


func (h *userHandler) GetUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	user, err := h.s.FindByID(ctx, userID)
	if err != nil {
		ctx.Error(err)
	}
	
	ctx.JSON(http.StatusOK, user)
}
func (h *userHandler) CreateUser(ctx *gin.Context){
	req := &UserCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	userID, err := h.s.RegisterUser(ctx, req.Username)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": userID})
}

func (h *userHandler) DeleteUser(ctx *gin.Context){
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	err = h.s.DeRegisterUser(ctx, userID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewUserHandler(s Service) *userHandler {
	return &userHandler{s: s}
}