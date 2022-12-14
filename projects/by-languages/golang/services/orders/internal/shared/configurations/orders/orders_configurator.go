package orders

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	grpcServer "github.com/mehdihadeli/store-golang-microservice-sample/pkg/grpc"
	customEcho "github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/custom_echo"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/configurations"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/shared/configurations/infrastructure"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/shared/web"
	"net/http"
)

type OrdersServiceConfigurator interface {
	ConfigureProductsModule() error
}

type ordersServiceConfigurator struct {
	ehcoServer customEcho.EchoHttpServer
	grpcServer grpcServer.GrpcServer
	*infrastructure.InfrastructureConfiguration
}

func NewOrdersServiceConfigurator(infrastructureConfiguration *infrastructure.InfrastructureConfiguration, ehcoServer customEcho.EchoHttpServer, grpcServer grpcServer.GrpcServer) *ordersServiceConfigurator {
	return &ordersServiceConfigurator{InfrastructureConfiguration: infrastructureConfiguration, ehcoServer: ehcoServer, grpcServer: grpcServer}
}

func (c *ordersServiceConfigurator) ConfigureOrdersService(ctx context.Context) error {
	pc := configurations.NewOrdersModuleConfigurator(c.InfrastructureConfiguration)
	err := pc.ConfigureOrdersModule(ctx)
	if err != nil {
		return err
	}

	err = c.migrateOrders()
	if err != nil {
		return err
	}

	c.ehcoServer.GetEchoInstance().GET("", func(ec echo.Context) error {
		return ec.String(http.StatusOK, fmt.Sprintf("%s is running...", web.GetMicroserviceName(c.Cfg)))
	})

	return nil
}
