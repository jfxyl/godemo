package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	common "godemo/etcd"
	"time"
)

func main() {
	var (
		err        error
		ctx        context.Context
		watchEvent *clientv3.Event
		watchResp  clientv3.WatchResponse
		watchChan  clientv3.WatchChan
		putResp    *clientv3.PutResponse
	)
	ctx = context.TODO()

	//PUT数据并返回PUT前数据
	go func() {
		for {
			if putResp, err = common.Client.Put(ctx, "/demo/key4", "value4", clientv3.WithPrevKV()); err != nil {
				panic(err)
			}
			if putResp.PrevKv != nil {
				fmt.Println("putPrevKv：", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
			} else {
				fmt.Println("putPrevKv：", putResp.PrevKv)
			}
			if _, err = common.Client.Delete(ctx, "/demo/key4"); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	watchChan = common.Client.Watch(ctx, "/demo", clientv3.WithPrefix(), clientv3.WithPrevKV())
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
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Minute)
}
