package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/hapiio/hapiio/core"
	"github.com/hapiio/hapiio/db"
	table "github.com/hapiio/hapiio/proto/table/v1"
	"github.com/hapiio/hapiio/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// type server struct {
// 	table.UnimplementedTableServiceServer
// }

// func NewServer() *server {
// 	return &server{}
// }

// func (s *server) SayHello(ctx context.Context, in *table.CreateTableRequest) (*table.CreateTableResponse, error) {
// 	return &table.CreateTableResponse{
// 		Table: &table.Table{
// 			TableName: "test",
// 			Column:    make([]*table.Column, 0),
// 		},
// 	}, nil
// }

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
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:8000"
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	table.RegisterTableServiceServer(s, server.New())
	// Serve gRPC Server
	log.Info("Serving gRPC on %s", addr)
	// go func() {
	log.Fatal(s.Serve(lis))
	// }()
	// err = gateway.Run("dns:///" + addr)
	// log.Fatalln(err)

}
