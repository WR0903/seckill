## 秒杀

挺羡慕别人写高并发的东西，最近看了一本书，书上恰好有一个秒杀的项目，自己模仿着构建一下。

## 启动方法

### gateway启动方法

```docker run -d -p 9411:9411 openzipkin/zipkin```运行zipkin；  
```consul agent -dev```启动consul   

进入gateway目录，运行```go run main.go```