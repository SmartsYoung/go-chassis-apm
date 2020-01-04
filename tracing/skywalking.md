# Skywalking

go-chassis-apm is a plugin of go chassis, it reports tracing data to skywalking server

## Configurations
1.import go-chassis-apm plugin pkg in main.go
```go
import _ "github.com/go-chassis/go-chassis-apm/tracing/skywalking"
```
**then, Calling function Init() in middleware/apm/apm_manager.go manually to init skywalking manager.**



2.edit conf/monitoring.yaml

**servicecomb.apm.tracing.tracer**
>  *(optional, string)* tracer'name, only skywalking now

**servicecomb.apm.tracing.settings.URI**
>  *(optional, string)* URI server address of skywalking

**servicecomb.apm.tracing.settings.servertype**
>  *(optional, string)* service type, match componentid in skywalking 
>  ex:  5001:servicecomb-mesher 5002:servicecomb-service-center 28:servicecomb-java-cahssis 

**servicecomb.apm.tracing.settings.enable**
>  *(optional, bool)* enable skywalking tracing ability

3.Add handler name which are defined in github.com/go-chassis/go-chassis-apm/handler

skywalking-consumer and skywalking-provider

## Example
```yaml
servicecomb:
  apm:                                #application performance monitor
    tracing:
      tracer: skywalking
      settings:
        enable: true                  #enable tracing ability
        URI: 127.0.0.1:11800          #url of skywalking 
        serverType: 5001              #server type
```

```
handler:
  chain:
    Provider:
      default: skywalking-provider
    Consumer:
      default: router,loadbalance,skywalking-consumer,transport
```

# 1. SkyWawlking-Manager Init
**You must init skywawlking manager pkg which will manage connection and report msg to skywalking**
- For example:
- [1] You can import skywalking manager proxy/pkg/skywalking in file proxy/bootstrap/bootstrap.go.
- [2] Calling function Init() in proxy/pkg/skywalking manually to init skywalking manager.
- [3] Adding skywalking's consumer handler name SkyWalkingConsumer defined in proxy/pkg/skywalking to consumerChain.
- [4] Adding skywalking's provider handler name SkyWalkingProvider defined in proxy/pkg/skywalking to providerChain.
- more details about handler chains in [go-chassis](https://github.com/go-chassis/go-chassis#readme)

# 2. SkyWalking-Handler Init
- You must import proxy/handler pkg to init skywalking handler. Not only skywalking handler, all the handlers which are customized for mesher are defined here.
- For example you can import handler pkg in file cmd/mesher/mesher.go
