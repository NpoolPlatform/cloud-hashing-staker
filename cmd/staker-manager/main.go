package main

import (
	"fmt"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/staker-manager/pkg/servicename"

	archivementconst "github.com/NpoolPlatform/archivement-manager/pkg/message/const"

	chainconst "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	ledgerconst "github.com/NpoolPlatform/ledger-manager/pkg/message/const"
	miningconst "github.com/NpoolPlatform/mining-manager/pkg/message/const"
	orderconst "github.com/NpoolPlatform/order-middleware/pkg/message/const"

	cli "github.com/urfave/cli/v2"
)

const serviceName = servicename.ServiceName

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		serviceName, serviceName)
	err := app.Init(
		serviceName,
		description,
		"",
		"",
		"./",
		nil,
		commands,
		ledgerconst.ServiceName,
		miningconst.ServiceName,
		archivementconst.ServiceName,
		orderconst.ServiceName,
		chainconst.ServiceName,
	)
	if err != nil {
		logger.Sugar().Errorf("fail to create %v: %v", serviceName, err)
		return
	}
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorf("fail to run %v: %v", serviceName, err)
	}
}
