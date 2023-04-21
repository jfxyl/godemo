package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	common "godemo/opentelemetry"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		err   error
		app   *App
		sigCh = make(chan os.Signal, 1)
		errCh = make(chan error)
		tp    *trace.TracerProvider
		cf    common.CloseFunc
	)
	//使用文件作为Exporter
	//if tp, cf, err = common.TracerProviderByFile(common.SERVICE_FIB); err != nil {
	//	panic(err)
	//}
	//使用jaeger作为Exporter
	if tp, cf, err = common.TracerProviderByJaeger(common.SERVICE_FIB, common.JaegerURL); err != nil {
		panic(err)
	}
	defer cf()

	//设置全局TracerProvider
	otel.SetTracerProvider(tp)

	signal.Notify(sigCh, os.Interrupt)
	app = NewApp(os.Stdin, log.New(os.Stdout, "", 0))
	go func() {
		errCh <- app.Run(context.TODO())
	}()

	select {
	case <-sigCh:
		app.l.Println("\n再见")
		return
	case err := <-errCh:
		app.l.Println(err)
	}
	time.Sleep(5 * time.Second)
}
