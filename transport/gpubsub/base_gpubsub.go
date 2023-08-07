package gpubsub

import (
	"context"
	"log"
	"sync"
	"template/contract"

	"cloud.google.com/go/pubsub"
)

type SubsConfig struct {
	SubscriptionID string
	TopicName      string
}

type templateGPubSubServer struct {
	Context         context.Context
	Client          *pubsub.Client
	SubsConfigSlice []SubsConfig
}

func getTemplateGPubSubServer(ctx context.Context, client *pubsub.Client, cfgs []SubsConfig) contract.TemplateServiceGPubSubInterface {
	return &templateGPubSubServer{
		Context:         ctx,
		Client:          client,
		SubsConfigSlice: cfgs,
	}
}

func NewTemplateGPubSubServer(projectID string) (*templateGPubSubServer, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return getTemplateGPubSubServer(ctx, client, getSubsConfigMaps()).(*templateGPubSubServer), nil
}

func NewTemplateGPubSubServerWithSubs(projectID string, subs []SubsConfig) (*templateGPubSubServer, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return getTemplateGPubSubServer(ctx, client, subs).(*templateGPubSubServer), nil
}

func RunSubscriber(s *templateGPubSubServer) {
	var wg sync.WaitGroup
	for _, v := range s.SubsConfigSlice {
		wg.Add(1)
		go subscribeTopic(s, v, &wg)
	}
	wg.Wait()
}

func subscribeTopic(s *templateGPubSubServer, cfg SubsConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	sub := s.Client.Subscription(cfg.SubscriptionID)
	exists, err := sub.Exists(s.Context)
	if err != nil {
		log.Println(err)
		return
	}

	if !exists {
		topic := s.Client.Topic(cfg.TopicName)
		exists, err := topic.Exists(s.Context)
		if err != nil {
			log.Println(err)
			return
		}
		if !exists {
			topic, err = s.Client.CreateTopic(s.Context, cfg.TopicName)
			if err != nil {
				log.Println(err)
				return
			}
		}

		sub, err = s.Client.CreateSubscription(
			s.Context,
			cfg.SubscriptionID,
			pubsub.SubscriptionConfig{
				Topic: topic,
			},
		)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if err = handleMessage(s, sub, cfg.TopicName); err != nil {
		log.Println(err)
		return
	}
}

func handleMessage(s *templateGPubSubServer, sub *pubsub.Subscription, topic string) error {
	switch topic {
	case "sample-add":
		return sub.Receive(s.Context, s.Add)
	}
	return nil
}

func getSubsConfigMaps() []SubsConfig {
	return []SubsConfig{
		{
			SubscriptionID: "add.sample-subs",
			TopicName:      "sample-add",
		},
	}
}
