package auth

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Handler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Signup(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type authHandler struct {
	s Service 
}

func (h *authHandler) Login(ctx *gin.Context){
	req := &LoginReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	accessToken, refreshToken, userID, err := h.s.LoginUser(ctx, req.Username)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie("access_token", accessToken, 15 * 60, "/", "localhost", false, true)

	ctx.SetCookie("refresh_token", refreshToken, 12 * 60 * 60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"user_id": userID, "username": req.Username})
}
func (h *authHandler) Logout(ctx *gin.Context){
	//TODO: change domain
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.Status(http.StatusNoContent)
}

func (h *authHandler) Signup(ctx *gin.Context){
	req := &SignupReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	}
	_, err := h.s.RegisterUser(ctx, req.Username)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"username": req.Username})
}

func (h *authHandler) Refresh(ctx *gin.Context){
	token, err := ctx.Cookie("refresh_token")
	if err != nil {
			ctx.Error(generalErrors.ErrUnauthorized)
			return
	}
	claim, err := h.s.ValidateRefreshToken(token)
	if err != nil {
		switch {
			case errors.Is(err, ErrTokenMalformed):
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token malformed"})
			case errors.Is(err, ErrTokenSignatureInvalid):
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token signature is invalid"})
			case errors.Is(err, ErrTokenExpired) || errors.Is(err, ErrTokenNotValidYet):
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		}
		return
	}
	userID, err := uuid.Parse(claim.UserID)
	if err != nil {
		ctx.Error(generalErrors.ErrInvalid)
		return
	} 
	accessToken, err := h.s.GenAccessToken(userID)

	ctx.SetCookie("access_token", accessToken, 15 * 60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"user_id": userID})
}

func NewAuthHandler(s Service) *authHandler{
	return &authHandler{s: s}
}