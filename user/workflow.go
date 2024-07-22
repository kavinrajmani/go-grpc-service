package user

import "context"

type Hdlr struct {
	UnimplementedUserServer
}

func NewHdlr() *Hdlr {
	return &Hdlr{}
}

func (s Hdlr) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return &UserResponse{
		Id:    req.Id,
		Name:  "Devdarshan Kavinraj",
		Email: "devdarshan@gmail.com",
	}, nil
}
