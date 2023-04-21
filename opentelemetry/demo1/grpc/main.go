package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	pb "godemo/grpc/proto/person"
	common "godemo/opentelemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

var (
	server *grpc.Server
	client pb.PersonServiceClient
)

func AuthUnaryInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("info.FullMethod", info.FullMethod)
	check := func() (err error) {
		md, ok := metadata.FromIncomingContext(c)
		if !ok {
			return fmt.Errorf("no auth info")
		}
		if len(md["authorization"]) == 0 {
			return fmt.Errorf("no auth info")
		}
		if md["authorization"][0] != "token123" {
			return fmt.Errorf("invalid token")
		}
		return nil
	}
	if info.FullMethod == "/person.PersonService/GetPerson" {
		if err := check(); err != nil {
			//return nil, err
		}
	}
	return handler(c, req)
}

type PersonService struct {
	pb.UnimplementedPersonServiceServer
}

func (p *PersonService) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.Person, error) {
	var (
		now       time.Time = time.Now()
		person    *pb.Person
		jsonBytes []byte
	)
	_, span := otel.Tracer(common.TRACER_PERSON).Start(ctx, "grpc.getPerson")
	defer span.End()
	person = &pb.Person{
		Id:       1,
		Name:     "jfxy",
		Gender:   1,
		Birthday: "1994-12-26",
		Avatar:   "",
		Email:    "",
		Phone:    "135********",
		Address: map[string]string{
			"province": "江苏省",
			"city":     "南京市",
			"district": "玄武区",
		},
		Tags: []string{},
		//Tags: []string{"golang", "php"},
		CreatedAt: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
		//CreatedAt:timestamppb.Now(),
	}
	jsonBytes, _ = json.Marshal(person)
	fmt.Println(string(jsonBytes))
	return person, nil
}

func main() {
	var (
		err error
		lis net.Listener
		tp  *sdktrace.TracerProvider
		cf  common.CloseFunc
	)

	//使用jaeger作为Exporter
	if tp, cf, err = common.TracerProviderByJaeger(common.SERVICE_PERSON, common.JaegerURL); err != nil {
		panic(err)
	}
	defer cf()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	//启动服务端
	if lis, err = net.Listen("tcp", common.GrpcPort); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			AuthUnaryInterceptor,
		),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	pb.RegisterPersonServiceServer(server, &PersonService{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
