package main

import (
	"bufio"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"strconv"
)

const (
	name = "fib"
)

type App struct {
	r io.Reader
	l *log.Logger
}

func (a *App) Run(ctx context.Context) error {
	for {
		newCtx, span := otel.Tracer(name).Start(ctx, "Run")
		n, err := a.Poll(newCtx)
		if err != nil {
			span.End()
			return err
		}
		a.Write(newCtx, n)
		span.End()
	}
}

func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(name).Start(ctx, "Poll")
	defer span.End()
	a.l.Print("What Fibonacci number would you like to know:")
	reader := bufio.NewReader(a.r)
	bytes, _, err := reader.ReadLine()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}
	str := string(bytes)
	span.SetAttributes(attribute.String("input", str))
	n, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return uint(n), err
}

func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = otel.Tracer(name).Start(ctx, "Write")
	defer span.End()

	f, err := func(ctx context.Context) (uint64, error) {
		_, span := otel.Tracer(name).Start(ctx, "Fib")
		defer span.End()
		return Fibonacci(n)
	}(ctx)
	if err != nil {
		a.l.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}

func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r, l}
}
