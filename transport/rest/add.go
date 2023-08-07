package rest

import (
	"io"
	"net/http"
	"template/contract"
	"template/model"
	"template/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

func (s *templateRestServer) Add(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var parsed contract.AddRequest
	err = protojson.Unmarshal(reqBody, &parsed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	req := &model.AddRequest{
		Param1: parsed.GetFirst(),
		Param2: parsed.GetSecond(),
	}
	res, err := usecase.Add(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	out := &contract.AddResponse{
		Result: res.Response,
	}
	resBody, err := protojson.Marshal(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(resBody)
}
