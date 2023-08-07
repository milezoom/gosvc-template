package gpubsub

import (
	"context"
	"log"
	"template/contract"
	"template/model"
	"template/usecase"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *templateGPubSubServer) Add(ctx context.Context, m *pubsub.Message) {
	var parsed contract.AddRequest
	err := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}.Unmarshal(m.Data, &parsed)
	if err != nil {
		log.Println(err)
		return
	}
	defer m.Ack()

	req := &model.AddRequest{
		Param1: parsed.GetFirst(),
		Param2: parsed.GetSecond(),
	}
	res, err := usecase.Add(req)
	if err != nil {
		log.Println(err)
		return
	}

	out := &contract.AddResponse{
		Result: res.Response,
	}
	resBody, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(out)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(resBody))
}
