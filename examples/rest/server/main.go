package main

import (
	"github.com/go-chassis/go-chassis"
	_ "github.com/go-chassis/go-chassis-apm/handler"
	"github.com/go-chassis/go-chassis-apm/middleware/apm"
	_ "github.com/go-chassis/go-chassis-apm/tracing/skywalking"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/core/server"
	"github.com/go-chassis/go-chassis/examples/schemas"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/

func main() {
	chassis.RegisterSchema("rest", &schemas.RestFulHello{}, server.WithSchemaID("RestHelloService"))
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	if err := apm.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
