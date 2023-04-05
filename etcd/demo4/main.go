package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	common "godemo/etcd"
	"time"
)

func main() {
	var (
		err                    error
		ctx                    context.Context = context.TODO()
		cancelFunc             context.CancelFunc
		leaseGrantResp         *clientv3.LeaseGrantResponse
		leaseKeepAliveResp     *clientv3.LeaseKeepAliveResponse
		leaseKeepAliveRespChan <-chan *clientv3.LeaseKeepAliveResponse
		txn                    clientv3.Txn
		txnResp                *clientv3.TxnResponse
	)
	ctx = context.TODO()
	//获取一个10s的租约
	if leaseGrantResp, err = common.Client.Grant(ctx, 10); err != nil {
		panic(err)
	}
	//自动续租
	ctx, cancelFunc = context.WithCancel(ctx)
	defer cancelFunc()
	defer common.Client.Revoke(context.TODO(), leaseGrantResp.ID)
	if leaseKeepAliveRespChan, err = common.Client.KeepAlive(ctx, leaseGrantResp.ID); err != nil {
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
					fmt.Println("续约成功 leaseKeepAliveResp.ID：", leaseGrantResp.ID)
				}
			}
			time.Sleep(1 * time.Second)
		}
	END:
	}()

	//获取一个锁
	//如果一个键不存在，获取它的create_revision时就是0
	txn = common.Client.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision("/demo/lock1"), "=", 0)).
		Then(clientv3.OpPut("/demo/lock1", "1", clientv3.WithLease(leaseGrantResp.ID))).
		Else(clientv3.OpGet("/demo/lock1"), clientv3.OpGet("/demo/lock1"))
	if txnResp, err = txn.Commit(); err != nil {
		panic(err)
	}
	if !txnResp.Succeeded {
		fmt.Println("锁被占用：", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}
	fmt.Println("处理业务开始")
	time.Sleep(10 * time.Second)
	fmt.Println("处理业务结束")

	if _, err = common.Client.Delete(ctx, "/demo/lock1"); err != nil {
		panic(err)
	}
}
