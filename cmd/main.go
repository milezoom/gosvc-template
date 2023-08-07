package main

import (
	"context"
	"log"
	"template/contract"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := contract.NewTemplateServiceClient(conn)
	request := &contract.AddRequest{
		First:  1,
		Second: 2,
	}
	var header metadata.MD
	response, err := client.Add(ctx, request, grpc.Header(&header))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response.String())
	log.Println(header["hello"][0])
}
