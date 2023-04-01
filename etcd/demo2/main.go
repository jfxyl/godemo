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
		err        error
		client     *clientv3.Client
		ctx        context.Context
		watchEvent *clientv3.Event
		watchResp  clientv3.WatchResponse
		watchChan  clientv3.WatchChan
		putResp    *clientv3.PutResponse
	)
	ctx = context.TODO()
	if client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}); err != nil {
		panic(err)
	}

	//PUT数据并返回PUT前数据
	go func() {
		for {
			if putResp, err = client.Put(ctx, "/demo/key4", "value4", clientv3.WithPrevKV()); err != nil {
				panic(err)
			}
			if putResp.PrevKv != nil {
				fmt.Println("putPrevKv：", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
			} else {
				fmt.Println("putPrevKv：", putResp.PrevKv)
			}
			if _, err = client.Delete(ctx, "/demo/key4"); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	watchChan = client.Watch(ctx, "/demo", clientv3.WithPrefix(), clientv3.WithPrevKV())
	go func() {
		fmt.Println("len(watchChan)", len(watchChan))
		//for watchResp = range watchChan {
		//	for _, watchEvent = range watchResp.Events {
		//		switch watchEvent.Type {
		//		case mvccpb.PUT:
		//			fmt.Println("PUT")
		//			if watchEvent.PrevKv != nil {
		//				fmt.Println(string(watchEvent.PrevKv.Key), string(watchEvent.PrevKv.Value))
		//			}
		//			if watchEvent.Kv != nil {
		//				fmt.Println(string(watchEvent.Kv.Key), string(watchEvent.Kv.Value))
		//			}
		//		case mvccpb.DELETE:
		//			fmt.Println("DELETE")
		//			if watchEvent.PrevKv != nil {
		//				fmt.Println(string(watchEvent.PrevKv.Key), string(watchEvent.PrevKv.Value))
		//			}
		//			if watchEvent.Kv != nil {
		//				fmt.Println(string(watchEvent.Kv.Key), string(watchEvent.Kv.Value))
		//			}
		//		}
		//	}
		//}

		for {
			select {
			case watchResp = <-watchChan:
				for _, watchEvent = range watchResp.Events {
					switch watchEvent.Type {
					case mvccpb.PUT:
						fmt.Println("PUT")
						if watchEvent.PrevKv != nil {
							fmt.Println(string(watchEvent.PrevKv.Key), string(watchEvent.PrevKv.Value))
						}
						if watchEvent.Kv != nil {
							fmt.Println(string(watchEvent.Kv.Key), string(watchEvent.Kv.Value))
						}
					case mvccpb.DELETE:
						fmt.Println("DELETE")
						if watchEvent.PrevKv != nil {
							fmt.Println(string(watchEvent.PrevKv.Key), string(watchEvent.PrevKv.Value))
						}
						if watchEvent.Kv != nil {
							fmt.Println(string(watchEvent.Kv.Key), string(watchEvent.Kv.Value))
						}
					}
				}
			default:
				fmt.Println("通道没有数据")
			}
		}
	}()

	time.Sleep(1 * time.Hour)
}
