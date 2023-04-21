package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	pb "godemo/grpc/proto/person"
	common "godemo/opentelemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const (
	TRACER_NAME = "person"
)

var (
	client pb.PersonServiceClient
)

func main() {
	var (
		err    error
		tp     *sdktrace.TracerProvider
		cf     common.CloseFunc
		conn   *grpc.ClientConn
		router *mux.Router
	)
	//使用jaeger作为Exporter
	if tp, cf, err = common.TracerProviderByJaeger(common.SERVICE_PERSON, common.JaegerURL); err != nil {
		panic(err)
	}
	defer cf()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	//启动客户端
	if conn, err = grpc.Dial(common.GrpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	); err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewPersonServiceClient(conn)

	//定义路由
	router = mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := otel.Tracer(common.TRACER_PERSON).Start(r.Context(), "http.middleware")
			defer span.End()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer(common.TRACER_PERSON).Start(r.Context(), "http.getPerson")
		defer span.End()
		person, err := client.GetPerson(ctx, &pb.GetPersonRequest{})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error：%s", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		jsonBytes, _ := json.Marshal(person)
		fmt.Fprintf(w, string(jsonBytes))
	})
	if err = http.ListenAndServe(":1111", router); err != nil {
		panic(err)
	}
}
