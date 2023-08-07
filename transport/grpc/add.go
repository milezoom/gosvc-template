package grpc

import (
	"context"
	"log"
	"template/contract"
	"template/model"
	"template/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *templateGrpcServer) Add(ctx context.Context, input *contract.AddRequest) (*contract.AddResponse, error) {
	req := &model.AddRequest{
		Param1: input.GetFirst(),
		Param2: input.GetSecond(),
	}
	res, err := usecase.Add(req)
	if err != nil {
		return nil, err
	}

	out := &contract.AddResponse{
		Result: res.Response,
	}

	header := metadata.New(map[string]string{
		"hello": "world",
	})
	if err := grpc.SendHeader(ctx, header); err != nil {
		log.Println(err)
	}

	return out, nil
}
