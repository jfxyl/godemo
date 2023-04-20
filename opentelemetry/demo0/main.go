package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	common "godemo/opentelemetry"
	"log"
	"os"
	"os/signal"
	"time"
)

type closeFunc func()

func main() {
	var (
		err   error
		app   *App
		sigCh = make(chan os.Signal, 1)
		errCh = make(chan error)
		tp    *trace.TracerProvider
		cf    closeFunc
	)
	//使用文件作为Exporter
	//if tp, ch, err = tracerProviderByFile(); err != nil {
	//	panic(err)
	//}
	//使用jaeger作为Exporter
	if tp, cf, err = tracerProviderByJaeger("http://localhost:14268/api/traces"); err != nil {
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

func tracerProviderByFile() (tp *trace.TracerProvider, cf closeFunc, err error) {
	var (
		file *os.File
		exp  trace.SpanExporter
	)
	file, err = os.OpenFile(common.FilePath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, func() {}, err
	}
	cf = func() {
		fmt.Println(file)
		file.Close()
	}
	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	return tp, cf, nil
}

func tracerProviderByJaeger(url string) (tp *trace.TracerProvider, cf closeFunc, err error) {
	var (
		exp trace.SpanExporter
	)
	exp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, func() {}, err
	}
	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	return tp, func() {}, nil
}

func newResource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("fib"),
		semconv.ServiceVersion("v0.1.0"),
		attribute.String("environment", "demo"),
	)
}
