package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeAddEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Add)
		res := s.TestAdd(ctx, req)
		return res, nil
	}
}

type EndPointServer struct {
	AddEndPoint endpoint.Endpoint
}

func NewEndPointServer(svc Service) EndPointServer {
	var addEndPoint endpoint.Endpoint
	{
		addEndPoint = MakeAddEndPoint(svc)
	}
	return EndPointServer{AddEndPoint: addEndPoint}
}
func (s EndPointServer) Add(ctx context.Context, in Add) AddAck {
	res, _ := s.AddEndPoint(ctx, in)
	return res.(AddAck)
}
