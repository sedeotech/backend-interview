package grpc

import (
	"context"

	"github.com/sedeotech/backend-interview/internal/services/cart"
	"github.com/sedeotech/backend-interview/pkg/rpc/cartpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartServer struct {
	cartpb.UnimplementedCartServer
	cartService cart.CartService
	//itemService     item.Service
}

func NewCartServer(c cart.CartService) *CartServer {
	return &CartServer{
		cartService: c,
	}
}

func (srv *CartServer) Get(ctx context.Context, req *cartpb.CartGetMsg) (*cartpb.CartGetResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (srv *CartServer) ConfigureAddress(ctx context.Context, req *cartpb.CartConfigureAddressMsg) (*cartpb.CartConfigureResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureAddress not implemented")
}
func (srv *CartServer) AddItem(ctx context.Context, req *cartpb.CartAddItemMsg) (*cartpb.CartAddItemResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (srv *CartServer) UpdateQuantity(ctx context.Context, req *cartpb.CartUpdateQuantityMsg) (*cartpb.CartUpdateQuantityResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuantity not implemented")
}
