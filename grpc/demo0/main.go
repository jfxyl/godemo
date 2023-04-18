package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "godemo/grpc/proto/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

type AuthInterceptor struct{}

func (a *AuthInterceptor) UnaryInterceptorfunc(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
			return nil, err
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

func (p *PersonService) ListPerson(ctx context.Context, req *pb.ListPersonRequest) (*pb.ListPersonResponse, error) {
	var (
		now    time.Time = time.Now()
		person *pb.Person
		resp   *pb.ListPersonResponse
	)
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

	resp = &pb.ListPersonResponse{
		Persons: []*pb.Person{
			person,
			person,
		},
	}
	return resp, nil
}

func main() {
	var (
		err        error
		port       = ":8080"
		conn       *grpc.ClientConn
		person     *pb.Person
		listPerson *pb.ListPersonResponse
		jsonBytes  []byte
	)
	//启动服务端
	go runServer(port)

	time.Sleep(1 * time.Second)

	//启动客户端
	if conn, err = grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewPersonServiceClient(conn)

	//调用GetPerson
	person, err = client.GetPerson(context.Background(), &pb.GetPersonRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		//proto.Marshal和proto.Unmarshal会导致空切片变成nil
		jsonBytes, err = json.Marshal(person)
		fmt.Println(string(jsonBytes))
	}
	//调用ListPerson
	listPerson, err = client.ListPerson(context.Background(), &pb.ListPersonRequest{})
	if err != nil {
		fmt.Println(err)
	} else {
		//proto.Marshal和proto.Unmarshal会导致空切片变成nil
		jsonBytes, err = json.Marshal(listPerson)
		fmt.Println(string(jsonBytes))
	}
}

func runServer(port string) {
	var (
		err error
		lis net.Listener
	)
	if lis, err = net.Listen("tcp", port); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server = grpc.NewServer(
		grpc.UnaryInterceptor((&AuthInterceptor{}).UnaryInterceptorfunc),
	)
	pb.RegisterPersonServiceServer(server, &PersonService{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
