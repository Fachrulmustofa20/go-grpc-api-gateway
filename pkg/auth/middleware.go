package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("authorization")

	if auth == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token := strings.Split(auth, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
