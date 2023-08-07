package grpc

import "template/contract"

type templateGrpcServer struct {
	contract.UnimplementedTemplateServiceServer
}

func NewTemplateGrpcServer() contract.TemplateServiceServer {
	return &templateGrpcServer{}
}
