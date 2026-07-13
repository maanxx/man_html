package pubsub

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/configs"
)

var (
	serviceName = "pubsub-service"
	version     = "1.0.0"
)

func newPubSubServiceCtx() gosctx.ServiceContext {
	return gosctx.NewServiceContext(
		gosctx.WithName(serviceName),
		gosctx.WithComponent(gosctx.NewAppLoggerDaily(configs.KeyLoggerDaily)),
	)
}

var PubSubCmd = &cobra.Command{
  Use:   "pubsub",
  Version: version,
  Short: "pubsub go clean architecture service context",
  Long: `pubsub go clean architecture service context trainning`,
  Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("----------- Server ----------")
	serviceCtx := newPubSubServiceCtx()

	logSrv := serviceCtx.MustGet(configs.KeyLoggerDaily).(gosctx.AppLoggerDaily).GetLogger("server")

	logSrv.Info("------------ start server --------------")

	if err := serviceCtx.Load(); err != nil{
		logSrv.Fatal(err)
	}


  },
}