package product

import "context"

type Hdlr struct {
	UnimplementedProductServer
}

func NewHdlr() *Hdlr {
	return &Hdlr{}
}

func (hdrl *Hdlr) CreateProduct(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {

	name := req.Name

	return &ProductResponse{
		Id:          "200",
		Name:        name,
		Description: req.Description,
	}, nil
}
