module user-service

go 1.14

require (
	github.com/go-kit/kit v0.11.0
	github.com/gohouse/gorose/v2 v2.1.2
	github.com/gorilla/mux v1.6.2
	github.com/juju/ratelimit v1.0.1
	github.com/longjoy/micro-go-book v0.0.0-20210706113218-203bfb2e5b19
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/viper v1.7.1
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6
	google.golang.org/grpc v1.39.0
	pkg v0.0.0
)

replace pkg => ../pkg

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
