package common

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"math/rand"
	"time"
)

type TransactionListener struct{}

func NewTransactionListener() *TransactionListener {
	return &TransactionListener{}
}

func (dl *TransactionListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())
	//返回一个非负的伪随机int值
	rint := r.Intn(3) + 1
	fmt.Println(rint)
	return primitive.LocalTransactionState(rint)
}

func (dl *TransactionListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	return primitive.CommitMessageState
}

var (
	Producer            rocketmq.Producer
	TransactionProducer rocketmq.TransactionProducer
	PushConsumer        rocketmq.PushConsumer
	PullConsumer        rocketmq.PullConsumer
)

func init() {
	var (
		err      error
		entpoint string = "192.168.2.5:9876"
	)
	rlog.SetLogLevel("error")
	Producer, err = rocketmq.NewProducer(
		producer.WithNameServer([]string{entpoint}),
		producer.WithRetry(2),
		producer.WithGroupName("ProducerGroup"),
		//根据msg设置的ShardingKey来选择写入到的queue
		producer.WithQueueSelector(producer.NewHashQueueSelector()),
	)
	if err != nil {
		panic(err)
	}

	TransactionProducer, err = rocketmq.NewTransactionProducer(
		NewTransactionListener(),
		producer.WithNameServer([]string{entpoint}),
		producer.WithRetry(1),
		producer.WithGroupName("TransactionProducerGroup"),
	)
	if err != nil {
		panic(err)
	}

	PushConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{entpoint}),
		consumer.WithGroupName("PushConsumerGroup"),
	)
	if err != nil {
		panic(err)
	}

	//2.1.1版本暂不支持pull consumer
	//PullConsumer, err = rocketmq.NewPullConsumer(
	//	consumer.WithNameServer([]string{entpoint}),
	//	consumer.WithGroupName("PullConsumerGroup"),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//err = PullConsumer.Subscribe("syncTopic", consumer.MessageSelector{})
	//if err = PullConsumer.Start(); err != nil {
	//	panic(err)
	//}
}
