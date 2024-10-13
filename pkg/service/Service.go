package service

import (
    "context"
)
type Service interface {
	TestAdd(ctx context.Context, in Add)AddAck

}
type Add struct {
	A int
	B int
}
type AddAck struct {
	Result int
}

type baseService struct {
}

func NewService() Service{
	return &baseService{}
}

func (s baseService) TestAdd(ctx context.Context, in Add)AddAck{
	return AddAck{Result: in.A + in.B}
}
