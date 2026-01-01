package user

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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


//TODO: add logging 
func (h *userHandler) GetUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} 
	user, err := h.s.FindByID(ctx, userID)
	if err != nil {
		slog.Error(err.Error())
		switch err {
			case ErrUserNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			case ErrUserExists:
				ctx.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
				return
			case ErrUncaught:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
		}
	}
	
	ctx.JSON(http.StatusOK, user)
}
func (h *userHandler) CreateUser(ctx *gin.Context){
	req := &UserCreateReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	userID, err := h.s.RegisterUser(ctx, req.Username)
	if err != nil {
		slog.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": userID})
}

func (h *userHandler) DeleteUser(ctx *gin.Context){
	userID, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = h.s.DeRegisterUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func NewUserHandler(s Service) *userHandler {
	return &userHandler{s: s}
}