package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	common "godemo/rocketmq"
	"time"
)

func main() {
	var (
		err error
		ctx context.Context
		msg *primitive.Message
		//resp      *primitive.SendResult
		//transResp *primitive.TransactionSendResult
	)
	if err = common.Producer.Start(); err != nil {
		panic(err)
	}
	if err = common.TransactionProducer.Start(); err != nil {
		panic(err)
	}
	err = common.PushConsumer.Subscribe("syncTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("msgs.length：%d\n", len(msgs))
		for _, msg := range msgs {
			fmt.Println(string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}
	if err = common.PushConsumer.Start(); err != nil {
		panic(err)
	}
	defer func() {
		common.Producer.Shutdown()
		common.TransactionProducer.Shutdown()
		common.PushConsumer.Shutdown()
	}()

	ctx = context.Background()
	//发送同步消息
	if _, err = common.Producer.SendSync(ctx, primitive.NewMessage("syncTopic", []byte("this is a test msg"))); err != nil {
		panic(err)
	}
	//发送异步消息
	err = common.Producer.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			panic(err)
		}
	}, primitive.NewMessage("asyncTopic", []byte("this is a test msg")))
	if err != nil {
		panic(err)
	}
	//发送单向消息
	if err = common.Producer.SendOneWay(ctx, primitive.NewMessage("oneWayTopic", []byte("this is a test msg"))); err != nil {
		panic(err)
	}
	//发送顺序消息
	//需要设置producer.WithQueueSelector(producer.NewHashQueueSelector()),通过给msg设置properties['SHARDING_KEY'],来选择写入的queue
	//且每次只能发送一条消息，源码中批量发送消息会重新包装msg,导致properties丢失
	shardingKey := []string{"a", "b", "c"}
	msg = &primitive.Message{
		Topic: "orderTopic",
		Body:  []byte("this is a test msg"),
	}
	for i := 0; i < 100; i++ {
		msg.WithShardingKey(shardingKey[i%len(shardingKey)])
		if _, err = common.Producer.SendSync(ctx, msg); err != nil {
			panic(err)
		}
	}
	//发送定时消息
	msg = &primitive.Message{
		Topic: "delayTopic",
		Body:  []byte("this is a test msg"),
	}
	if _, err = common.Producer.SendSync(ctx, msg.WithDelayTimeLevel(2)); err != nil {
		panic(err)
	}
	//发送事务消息
	msg = &primitive.Message{
		Topic: "transTopic",
		Body:  []byte("this is a test msg"),
	}
	if _, err = common.TransactionProducer.SendMessageInTransaction(ctx, msg); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
}
