package inx

import (
	"context"

	"go.uber.org/dig"

	"github.com/gohornet/inx-mqtt/pkg/daemon"
	"github.com/gohornet/inx-mqtt/pkg/nodebridge"
	"github.com/iotaledger/hive.go/app"
)

func init() {
	CoreComponent = &app.CoreComponent{
		Component: &app.Component{
			Name:     "INX",
			DepsFunc: func(cDeps dependencies) { deps = cDeps },
			Params:   params,
			Provide:  provide,
			Run:      run,
		},
	}
}

type dependencies struct {
	dig.In
	NodeBridge *nodebridge.NodeBridge
}

var (
	CoreComponent *app.CoreComponent
	deps          dependencies
)

func provide(c *dig.Container) error {
	return c.Provide(func() (*nodebridge.NodeBridge, error) {
		return nodebridge.NewNodeBridge(CoreComponent.Daemon().ContextStopped(),
			ParamsINX.Address,
			CoreComponent.Logger())
	})
}

func run() error {
	return CoreComponent.Daemon().BackgroundWorker("INX", func(ctx context.Context) {
		CoreComponent.LogInfo("Starting NodeBridge")
		deps.NodeBridge.Run(ctx)
		CoreComponent.LogInfo("Stopped NodeBridge")
	}, daemon.PriorityDisconnectINX)
}
