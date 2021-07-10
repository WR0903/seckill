module gateway

go 1.14

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-kit/kit v0.9.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/spf13/viper v1.7.1
	pb v0.0.0
	pkg v0.0.0
)

replace pkg => ../pkg

replace pb => ../pb

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
