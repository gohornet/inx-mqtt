package app

import (
	"github.com/gohornet/inx-mqtt/core/inx"
	"github.com/gohornet/inx-mqtt/core/mqtt"
	"github.com/gohornet/inx-mqtt/plugins/prometheus"
	"github.com/iotaledger/hive.go/app"
	"github.com/iotaledger/hive.go/app/plugins/profiling"
)

var (
	// Name of the app.
	Name = "inx-mqtt"

	// Version of the app.
	Version = "0.6.0"
)

func App() *app.App {
	return app.New(Name, Version,
		app.WithInitComponent(InitComponent),
		app.WithCoreComponents([]*app.CoreComponent{
			inx.CoreComponent,
			mqtt.CoreComponent,
		}...),
		app.WithPlugins([]*app.Plugin{
			profiling.Plugin,
			prometheus.Plugin,
		}...),
	)
}

var (
	InitComponent *app.InitComponent
)

func init() {
	InitComponent = &app.InitComponent{
		Component: &app.Component{
			Name: "App",
		},
		NonHiddenFlags: []string{
			"config",
			"help",
			"version",
		},
	}
}
