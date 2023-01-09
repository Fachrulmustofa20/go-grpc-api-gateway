package product

import (
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/auth"
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/config"
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreatedProduct)
	routes.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) CreatedProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
