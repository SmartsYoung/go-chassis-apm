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

package handler

import (
	"github.com/go-chassis/go-chassis-apm/apm"
	chassisHandler "github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/go-mesh/openlogging"
)

//SkyWalkingProviderHandler struct
type SkyWalkingProviderHandler struct {
}

//Handle is for provider
func (sp *SkyWalkingProviderHandler) Handle(chain *chassisHandler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	openlogging.Debug("SkyWalkingProviderHandler begin " + i.MicroServiceName)
	span, err := apm.CreateEntrySpan(i)
	if err != nil {
		openlogging.Error("CreateEntrySpan error: " + err.Error())
		return
	}
	chain.Next(i, func(r *invocation.Response) (err error) {
		err = cb(r)
		apm.EndSpan(span, r.Status)
		return
	})
}

//Name return provider name
func (sp *SkyWalkingProviderHandler) Name() string {
	return chassisHandler.SkyWalkingProvider
}

//NewSkyWalkingProvier return provider handler for SkyWalking
func newSkyWalkingProvider() chassisHandler.Handler {
	return &SkyWalkingProviderHandler{}
}

//SkyWalkingConsumerHandler struct
type SkyWalkingConsumerHandler struct {
}

//Handle is for consumer
func (sc *SkyWalkingConsumerHandler) Handle(chain *chassisHandler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	openlogging.Debug("SkyWalkingConsumerHandler begin " + i.MicroServiceName)
	span, err := apm.CreateEntrySpan(i)
	if err != nil {
		openlogging.Error("CreateEntrySpan error:" + err.Error())
		return
	}
	spanExit, err := apm.CreateExitSpan(i)
	if err != nil {
		openlogging.Error("CreateExitSpan error:" + err.Error())
		return
	}
	chain.Next(i, func(r *invocation.Response) (err error) {
		err = cb(r)
		apm.EndSpan(spanExit, r.Status)
		apm.EndSpan(span, r.Status)
		openlogging.Debug("SkyWalkingConsumerHandler end.")
		return
	})
}

//Name return consumer name
func (sc *SkyWalkingConsumerHandler) Name() string {
	return chassisHandler.SkyWalkingConsumer
}

//newSkyWalkingConsumer return consumer handler for SkyWalking
func newSkyWalkingConsumer() chassisHandler.Handler {
	return &SkyWalkingConsumerHandler{}
}

func init() {
	chassisHandler.HandlerFuncMap[chassisHandler.SkyWalkingConsumer] = newSkyWalkingConsumer
	chassisHandler.HandlerFuncMap[chassisHandler.SkyWalkingProvider] = newSkyWalkingProvider
}
