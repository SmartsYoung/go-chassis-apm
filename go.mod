module github.com/go-chassis/go-chassis-apm

require (
	github.com/SkyAPM/go2sky v0.2.0
	github.com/go-chassis/go-chassis v1.8.1
	github.com/go-chassis/paas-lager v1.1.0 // indirect
	github.com/go-mesh/openlogging v1.0.1
	github.com/stretchr/testify v1.4.0
)

//replace github.com/go-chassis/go-chassis latest => github.com/SmartsYoung/go-chassis v1.8.1-20200109142703-00724bde9095
replace github.com/go-chassis/go-chassis v1.8.1 => ../go-chassis

//replace github.com/go-chassis/go-chassis v1.8.1 => github.com/SmartsYoung/go-chassis v1.8.1

go 1.13
