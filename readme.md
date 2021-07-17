## 秒杀

挺羡慕别人写高并发的东西，最近看了一本书，书上恰好有一个秒杀的项目，自己模仿着构建一下。

## 启动方法

启动辅助工具：  
```docker run -d -p 9411:9411 openzipkin/zipkin```运行zipkin；  
```consul agent -dev```启动consul   
zookeeper启动 ```docker run --privileged=true -d --publish 2181:2181  -d zookeeper:latest```   
正常方式启动redis   

### gateway启动方法
进入gateway目录，运行```go run main.go```

### oauth-server启动方法
保持gateway启动，进入oauth-server目录，运行```go run main.go```


### sk-core启动方法
进入sk-core目录下，运行```go run main.go```

### sk-app启动方法
进入sk-app目录下，运行```go run main.go```

### sk-admin启动方法
进入sk-admin目录下，运行```go run main.go```

### user-service启动方法
进入user-servicez目录下，运行```go run main.go```
