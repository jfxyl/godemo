package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var (
		err    error
		client *clientv3.Client
		op     clientv3.Op
		ctx    context.Context
		kv     *mvccpb.KeyValue
		opResp clientv3.OpResponse
	)
	ctx = context.TODO()
	if client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}); err != nil {
		panic(err)
	}

	op = clientv3.OpPut("/demo/key5", "value5", clientv3.WithPrevKV())
	if opResp, err = client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Put().PrevKv)

	op = clientv3.OpPut("/demo/key6", "value6", clientv3.WithPrevKV())
	if opResp, err = client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Put().PrevKv)

	op = clientv3.OpGet("/demo", clientv3.WithPrefix())
	if opResp, err = client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Get().Kvs)

	op = clientv3.OpDelete("/demo", clientv3.WithPrevKV(), clientv3.WithPrefix())
	if opResp, err = client.Do(ctx, op); err != nil {
		panic(err)
	}

	for _, kv = range opResp.Del().PrevKvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}
}
