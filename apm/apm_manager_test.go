package apm

import (
	"context"
	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/config/model"
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/stretchr/testify/assert"
	"testing"
)

//initConfig
func initConfig() {
	config.MonitorCfgDef = &model.MonitorCfg{ServiceComb: model.ServiceCombStruct{APM: model.APMStruct{Tracing: model.TracingStruct{Settings: map[string]string{"URI": "127.0.0.1:11800", "enable": "true"}}}}}
	config.MicroserviceDefinition = &model.MicroserviceCfg{ServiceDescription: model.MicServiceStruct{Name: "skywalking"}}
}

//initInv
func initInv() *invocation.Invocation {
	i := invocation.New(context.Background())
	i.MicroServiceName = "test"
	i.Endpoint = "calculator"
	i.URLPathFormat = "/bmi"
	i.SetHeader("Sw6", "")
	return i
}

//TestCreateEntrySpan
func TestCreateEntrySpan(t *testing.T) {
	initConfig()
	err := Init()
	assert.Equal(t, err, nil)
	span, err := CreateEntrySpan(initInv())
	assert.Equal(t, err, nil)
	EndSpan(span, 1)
}

//TestExitSpan
func TestExitSpan(t *testing.T) {
	initConfig()
	span, err := CreateEntrySpan(initInv())
	assert.Equal(t, err, nil)
	EndSpan(span, 1)
}

//TestEndSpan
func TestEndSpan(t *testing.T) {
	initConfig()
	span, err := CreateEntrySpan(initInv())
	assert.Equal(t, err, nil)
	err = EndSpan(span, 1)
	assert.Equal(t, err, nil)
}

//TestInit
func TestInit(t *testing.T) {
	initConfig()
	err := Init()
	assert.Equal(t, err, nil)
}
