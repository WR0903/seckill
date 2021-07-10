module sk-core

go 1.14

require (
	github.com/go-kit/kit v0.11.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/spf13/viper v1.7.1
	pkg v0.0.0
)

replace pkg => ../pkg

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
