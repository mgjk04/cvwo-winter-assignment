package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/auth"
)

func AuthMiddleware(s auth.Service) gin.HandlerFunc{
	return func (ctx *gin.Context){
		token, err := ctx.Cookie("access_token")
		if err != nil {
			switch {
				case errors.Is(err, http.ErrNoCookie):
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"access token not found"})
				default:
					slog.Error(err.Error())
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
				}
				return
		}
		claim, err := s.ValidateAccessToken(token)
		if err != nil {
			switch {
				case errors.Is(err, auth.ErrTokenMalformed):
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"token malformed"})
				case errors.Is(err, auth.ErrTokenSignatureInvalid):
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"token signature is invalid"})
				case errors.Is(err, auth.ErrTokenExpired) || errors.Is(err, auth.ErrTokenNotValidYet):
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"token invalid"})
			}
			return
		}
		
		if claim != nil {
			ctx.Set("user_id", claim.UserID)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"token signature is invalid"})
		}
	}
}