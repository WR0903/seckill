package setup

import (
	"fmt"
	conf "pkg/config"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

//初始化Etcd
func InitZk() {
	var hosts = []string{"127.0.0.1:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	conf.Zk.ZkConn = conn
	conf.Zk.SecProductKey = "/product"
}
