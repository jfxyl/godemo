package common

import (
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"os"
)

type CloseFunc func()

var (
	DirPath  = "./opentelemetry/tmp/"
	FilePath = DirPath + "fib.txt"
)

func init() {
	var (
		err error
	)
	if _, err = os.Stat(DirPath); os.IsNotExist(err) {
		os.MkdirAll(DirPath, 0755)
	}
	if _, err = os.Stat(FilePath); os.IsNotExist(err) {
		if _, err = os.OpenFile(FilePath, os.O_CREATE, 0755); err != nil {
			panic(err)
		}
	}
}

func TracerProviderByFile() (tp *trace.TracerProvider, cf CloseFunc, err error) {
	var (
		file *os.File
		exp  trace.SpanExporter
	)
	file, err = os.OpenFile(FilePath, os.O_CREATE|os.O_WRONLY, 0777)
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

func TracerProviderByJaeger(url string) (tp *trace.TracerProvider, cf CloseFunc, err error) {
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
