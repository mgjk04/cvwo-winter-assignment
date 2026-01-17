package middleware

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)


//I think this can be simplified  into a map somehow, lots of repeated code
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			switch {
				case errors.Is(err, generalErrors.ErrNotFound):
					ctx.JSON(http.StatusNotFound, gin.H{"error":  err.Error()})
				case errors.Is(err, generalErrors.ErrConflict):
					ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				case errors.Is(err, generalErrors.ErrInvalid):
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				case errors.Is(err, generalErrors.ErrUnauthorized):
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				case errors.Is(err, generalErrors.ErrForbidden):
					ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				default:
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	}
}