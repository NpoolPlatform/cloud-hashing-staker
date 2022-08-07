package main

import (
	"github.com/NpoolPlatform/staker-manager/api"

	benefit "github.com/NpoolPlatform/staker-manager/pkg/benefit"
	migrate "github.com/NpoolPlatform/staker-manager/pkg/migrate"
	order "github.com/NpoolPlatform/staker-manager/pkg/order"
	collector "github.com/NpoolPlatform/staker-manager/pkg/sentinel/collector"
	limitation "github.com/NpoolPlatform/staker-manager/pkg/sentinel/limitation"
	withdraw "github.com/NpoolPlatform/staker-manager/pkg/sentinel/withdraw"
	transaction "github.com/NpoolPlatform/staker-manager/pkg/transaction"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := migrate.Migrate(c.Context); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()

		go transaction.Watch(c.Context)
		go order.Watch(c.Context)
		go collector.Watch(c.Context)
		go limitation.Watch(c.Context)
		go withdraw.Watch(c.Context)
		go benefit.Watch(c.Context)

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	apimgrcli.Register(mux)

	return nil
}
