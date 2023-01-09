package auth

import (
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
