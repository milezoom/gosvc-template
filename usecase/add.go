package usecase

import "template/model"

func Add(input *model.AddRequest) (*model.AddResponse, error) {
	return &model.AddResponse{
		Response: input.Param1 + input.Param2,
	}, nil
}
