package main

import (
	"log"
	"net"
	"os"
	"template/contract"

	tPubSub "template/transport/gpubsub"
	tGrpc "template/transport/grpc"
	tRest "template/transport/rest"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}
	mux := cmux.New(conn)

	gListen := mux.Match(cmux.HTTP2(), cmux.HTTP2HeaderField("content-type", "application/grpc"))
	gServer := grpc.NewServer()
	contract.RegisterTemplateServiceServer(gServer, tGrpc.NewTemplateGrpcServer())
	go gServer.Serve(gListen)

	rListen := mux.Match(cmux.HTTP1Fast())
	rServer := tRest.NewTemplateRestServer()
	go rServer.Serve(rListen)

	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8800")
	psServer, err := tPubSub.NewTemplateGPubSubServer("local-dev")
	if err != nil {
		panic(err)
	}
	go tPubSub.RunSubscriber(psServer)

	mux.Serve()
}
