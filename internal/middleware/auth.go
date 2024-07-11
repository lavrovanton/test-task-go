package middleware

import (
	"net/http"
	"test-task-go/internal/model"
	"test-task-go/pkg"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAdmin() (model.User, error)
}

func AuthMiddleware(repo UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Request.Header.Get("X-API-KEY")
		if len(apiKey) == 0 || len(apiKey) < 64 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": pkg.ErrAuthFailed.Error()})
			return
		}
		admin, err := repo.GetAdmin()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": pkg.ErrInternalServerError.Error()})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(admin.ApiKey), []byte(apiKey)) != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": pkg.ErrAuthFailed.Error()})
			return
		}
		ctx.Next()
	}
}
