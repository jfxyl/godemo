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
		err     error
		client  *clientv3.Client
		ctx     context.Context
		kv      *mvccpb.KeyValue
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
	)
	ctx = context.TODO()
	if client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}); err != nil {
		panic(err)
	}

	//PUT数据并返回PUT前数据
	if putResp, err = client.Put(ctx, "/demo/key1", "value1", clientv3.WithPrevKV()); err != nil {
		panic(err)
	}
	if putResp, err = client.Put(ctx, "/demo/key2", "value2", clientv3.WithPrevKV()); err != nil {
		panic(err)
	}
	if putResp.PrevKv != nil {
		fmt.Println("putPrevKv：", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
	} else {
		fmt.Println("putPrevKv：", putResp.PrevKv)
	}
	//GET以/demo为前缀的所有key的数据
	if getResp, err = client.Get(ctx, "/demo", clientv3.WithPrefix()); err != nil {
		panic(err)
	}
	if len(getResp.Kvs) > 0 {
		for _, kv = range getResp.Kvs {
			fmt.Println(string(kv.Key), string(kv.Value))
		}
	}
	//DELETE以/demo为前缀的所有key的数据并返回删除前的数据
	if delResp, err = client.Delete(ctx, "/demo", clientv3.WithPrefix(), clientv3.WithPrevKV()); err != nil {
		panic(err)
	}
	if len(delResp.PrevKvs) > 0 {
		for _, kv = range delResp.PrevKvs {
			fmt.Println("delPrevKv：", string(kv.Key), string(kv.Value))
		}
	} else {
		fmt.Println("delPrevKv：", delResp.PrevKvs)
	}
}
