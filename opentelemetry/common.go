package common

import (
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"os"
)

type CloseFunc func()

const (
	TRACER_PERSON  = "PERSON"
	TRACER_FIB     = "FIB"
	SERVICE_PERSON = "PERSON"
	SERVICE_FIB    = "FIB"
)

var (
	DirPath   = "./opentelemetry/tmp/"
	FilePath  = DirPath + "opentelemetry.txt"
	JaegerURL = "http://localhost:14268/api/traces"

	GrpcPort = ":7777"
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

func TracerProviderByFile(serviceName string) (tp *trace.TracerProvider, cf CloseFunc, err error) {
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
	exp, err = stdouttrace.New(
		stdouttrace.WithWriter(file),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		return nil, cf, err
	}
	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource(serviceName)),
	)
	return tp, cf, nil
}

func TracerProviderByJaeger(serviceName, url string) (tp *trace.TracerProvider, cf CloseFunc, err error) {
	var (
		exp trace.SpanExporter
	)
	exp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, func() {}, err
	}
	tp = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource(serviceName)),
	)
	return tp, func() {}, nil
}

func newResource(serviceName string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion("v0.1.0"),
		attribute.String("environment", "demo"),
	)
}
