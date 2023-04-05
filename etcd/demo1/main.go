package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	common "godemo/etcd"
	"sync"
	"time"
)

func main() {
	var (
		err                    error
		ctx                    context.Context
		timeoutctx             context.Context
		kv                     *mvccpb.KeyValue
		putResp                *clientv3.PutResponse
		getResp                *clientv3.GetResponse
		leaseGrantResp         *clientv3.LeaseGrantResponse
		leaseKeepAliveResp     *clientv3.LeaseKeepAliveResponse
		leaseKeepAliveRespChan <-chan *clientv3.LeaseKeepAliveResponse
		wg                     *sync.WaitGroup
	)
	ctx = context.TODO()
	wg = new(sync.WaitGroup)
	wg.Add(1)

	//获取一个10s的租约
	if leaseGrantResp, err = common.Client.Grant(ctx, 10); err != nil {
		panic(err)
	}
	//PUT数据携带上租约ID并返回PUT前数据
	if putResp, err = common.Client.Put(ctx, "/demo/key3", "value3", clientv3.WithPrevKV(), clientv3.WithLease(leaseGrantResp.ID)); err != nil {
		panic(err)
	}
	if putResp.PrevKv != nil {
		fmt.Println("putPrevKv：", string(putResp.PrevKv.Key), string(putResp.PrevKv.Value))
	} else {
		fmt.Println("putPrevKv：", putResp.PrevKv)
	}
	//自动续租5s
	timeoutctx, _ = context.WithTimeout(ctx, 5*time.Second)
	if leaseKeepAliveRespChan, err = common.Client.KeepAlive(timeoutctx, leaseGrantResp.ID); err != nil {
		panic(err)
	}
	//处理续租应答
	go func() {
		for {
			select {
			case leaseKeepAliveResp = <-leaseKeepAliveRespChan:
				if leaseKeepAliveResp == nil {
					fmt.Println("续约失败")
					goto END
				} else {
					fmt.Println("续约成功 leaseKeepAliveResp.ID：", leaseKeepAliveResp.ID)
				}
			}
			time.Sleep(1 * time.Second)
		}
	END:
	}()

	go func() {
		for {
			if getResp, err = common.Client.Get(ctx, "/demo/key3"); err != nil {
				panic(err)
			}
			if len(getResp.Kvs) > 0 {
				for _, kv = range getResp.Kvs {
					fmt.Println(string(kv.Key), string(kv.Value))
				}
			} else {
				wg.Done()
			}
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
