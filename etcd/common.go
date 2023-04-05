package common

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	Client *clientv3.Client
)

func init() {
	var (
		err error
	)
	if Client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}); err != nil {
		panic(err)
	}
}
