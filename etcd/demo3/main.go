package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	common "godemo/etcd"
)

func main() {
	var (
		err    error
		op     clientv3.Op
		ctx    context.Context
		kv     *mvccpb.KeyValue
		opResp clientv3.OpResponse
	)
	ctx = context.TODO()

	op = clientv3.OpPut("/demo/key5", "value5", clientv3.WithPrevKV())
	if opResp, err = common.Client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Put().PrevKv)

	op = clientv3.OpPut("/demo/key6", "value6", clientv3.WithPrevKV())
	if opResp, err = common.Client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Put().PrevKv)

	op = clientv3.OpGet("/demo", clientv3.WithPrefix())
	if opResp, err = common.Client.Do(ctx, op); err != nil {
		panic(err)
	}
	fmt.Println(opResp.Get().Kvs)

	op = clientv3.OpDelete("/demo", clientv3.WithPrevKV(), clientv3.WithPrefix())
	if opResp, err = common.Client.Do(ctx, op); err != nil {
		panic(err)
	}

	for _, kv = range opResp.Del().PrevKvs {
		fmt.Println(string(kv.Key), string(kv.Value))
	}
}
