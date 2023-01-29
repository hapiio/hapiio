package main

import (
	"io/ioutil"
	"net"
	"os"

	"github.com/hapiio/hapiio/cmd"
	"github.com/hapiio/hapiio/gateway"
	"github.com/hapiio/hapiio/insecure"
	table_pb "github.com/hapiio/hapiio/proto/table/v1"
	"github.com/hapiio/hapiio/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	table_pb.RegisterTableServiceServer(s, server.New())
	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
	cmd.Execute()
}
