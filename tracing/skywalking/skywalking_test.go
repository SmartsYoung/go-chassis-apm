/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package skywalking_test

import (
	"context"
	"github.com/go-chassis/go-chassis-apm/apm"
	"github.com/go-chassis/go-chassis-apm/tracing/skywalking"
	"github.com/stretchr/testify/assert"
	"testing"
)


var (
	op           apm.TracingOptions
	apmClient    apm.TracingClient
	sc           apm.SpanContext
)

func InitOption() {
	op = apm.TracingOptions{
		APMName:        "skywalking",
		ServerURI:      "192.168.88.64:8080",
		MicServiceName: "mesher",
		MicServiceType: 1}
}

func IniSpanContext() {
	sc = apm.SpanContext{
		Ctx:           context.Background(),
		OperationName: "test",
		ParTraceCtx:   map[string]string{},
		TraceCtx:      map[string]string{},
		Peer:          "test",
		Method:        "get",
		URL:           "/etc/url",
		ComponentID:   "1",
		SpanLayerID:   "11",
		ServiceName:   "mesher"}
}

func TestNewApmClient(t *testing.T) {
	InitOption()
	var err error
	apmClient, err = skywalking.NewApmClient(op)
	assert.Equal(t, err, nil)
}

func TestCreateEntrySpan(t *testing.T) {
	InitOption()
	IniSpanContext()
	var err error
	apmClient, err = skywalking.NewApmClient(op)
	assert.Equal(t, err, nil)
	span, err := apmClient.CreateEntrySpan(sc)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, span, nil)
}

func TestCreateExitSpan(t *testing.T) {
	InitOption()
	IniSpanContext()
	var err error
	apmClient, err = skywalking.NewApmClient(op)
	assert.Equal(t, err, nil)
	span, err := apmClient.CreateExitSpan(sc)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, span, nil)
}

func TestEndSpan(t *testing.T) {
	InitOption()
	IniSpanContext()
	var err error
	apmClient, err = skywalking.NewApmClient(op)
	assert.Equal(t, err, nil)
	span, err := apmClient.CreateEntrySpan(sc)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, span, nil)
	err = apmClient.EndSpan(span, 1)
	assert.Equal(t, err, nil)
}
