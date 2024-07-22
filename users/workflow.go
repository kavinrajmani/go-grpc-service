package users

import "context"

type Hdlr struct {
	UnimplementedUsersServer
}

func NewHdlr() *Hdlr {
	return &Hdlr{}
}

var users = []*User{
	{
		Id:    "1",
		Name:  "Devdarshan Kavinraj",
		Email: "devdarshan@gmail.com",
	},
	{
		Id:    "2",
		Name:  "Kavinraj Mani",
		Email: "kavinraj@gmail.com",
	},
	{
		Id:    "3",
		Name:  "Roobashri Kavinraj",
		Email: "roobashri@gmail.com",
	},
}

func (s Hdlr) GetUsers(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return &UserResponse{
		Users: users,
	}, nil
}

func (s Hdlr) CreateUser(ctx context.Context, req *User) (*CreateResponse, error) {
	user := &User{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}
	users = append(users, user)

	return &CreateResponse{
		Success: true,
		Message: "User created successfully",
	}, nil
}
