module github.com/go-chassis/go-chassis-apm

require (
	github.com/go-chassis/go-chassis v1.8.0
	github.com/go-chassis/paas-lager v1.1.0 // indirect
	github.com/go-mesh/openlogging v1.0.1
	github.com/stretchr/testify v1.4.0
	github.com/tetratelabs/go2sky v0.1.1-0.20190703154722-1eaab8035277
)

//github.com/go-chassis/go-chassis v1.8.0 => github.com/SmartsYoung/go-chassis v1.7.6-20191209201319-5e09cf871f4f
replace (
	github.com/go-chassis/go-chassis v1.8.0 => ../go-chassis
	github.com/tetratelabs/go2sky v0.1.1-0.20190703154722-1eaab8035277 => github.com/SkyAPM/go2sky v0.1.1-0.20190703154722-1eaab8035277
)

go 1.13
