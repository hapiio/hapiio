package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hapiio/hapiio/core"
	"github.com/hapiio/hapiio/db"
	table "github.com/hapiio/hapiio/proto/table/v1"
	"google.golang.org/grpc"
)

type server struct {
	table.UnimplementedTableServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *table.CreateTableRequest) (*table.CreateTableResponse, error) {
	return &table.CreateTableResponse{
		Table: &table.Table{
			TableName: "test",
			Column:    make([]*table.Column, 0),
		},
	}, nil
}

func main() {
	// provider, err := auth.Provider("facebook")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(provider.AuthUrl())

	// var cfg core.Config
	conf := core.Configure()
	fmt.Println(conf.Greeting)
	db.Connect(conf)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	table.RegisterTableServiceServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))

}
