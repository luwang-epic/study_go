package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/*
type Client struct {
    Cluster
    KV
    Lease
    Watcher
    Auth
    Maintenance
    // Username is a user name for authentication.
    Username string
    // Password is a password for authentication.
    Password string
    // contains filtered or unexported fields
}
类型中的成员是 etcd 客户端几何核心功能模块的具体实现，它们分别用于：
	Cluster：向集群里增加 etcd 服务端节点之类，属于管理员操作。
	KV：我们主要使用的功能，即 K-V 键值库的操作。
	Lease：租约相关操作，比如申请一个 TTL=10 秒的租约（应用给 key 可以实现键值的自动过期）。
	Watcher：观察订阅，从而监听最新的数据变化。
	Auth：管理 etcd 的用户和权限，属于管理员操作。
	Maintenance：维护 etcd，比如主动迁移 etcd 的 leader 节点，属于管理员操作。

我们通过方法 clientv3.NewKV() 来获得 KV 接口的实现（实现中内置了错误重试机制）：
	kv := clientv3.NewKV(cli)

 */
func getEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})
	return cli, err
}


func putKeyValue(key string, value string) {
	client, _ := getEtcdClient()
	kv := clientv3.NewKV(client)
	kv.Put(context.TODO(), key, value)
}

func getValue(key string) string {
	client, _ := getEtcdClient()
	kv := clientv3.NewKV(client)
	rep, _ := kv.Get(context.TODO(), key)
	return string(rep.Kvs[0].Value)
}


func main() {
	putKeyValue("name", "test")
	value := getValue("name")
	println("get name from etct is ", value)
}
