package product

import (
	"fmt"

	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/config"
	"github.com/Fachrulmustofa20/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(cfg *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect : ", err)
	}

	return pb.NewProductServiceClient(cc)
}
