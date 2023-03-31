package grpc

import (
	"context"

	"github.com/sedeotech/backend-interview/internal/services/item"
	"github.com/sedeotech/backend-interview/pkg/rpc/itempb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ItemServer struct {
	itempb.UnimplementedItemServer
	//cartService cart.CartService
	itemService item.ItemService
}

func NewItemServer(i item.ItemService) *ItemServer {
	return &ItemServer{
		itemService: i,
	}
}

func (srv *ItemServer) Create(ctx context.Context, req *itempb.ItemCreateMsg) (*itempb.ItemMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (srv *ItemServer) Update(ctx context.Context, req *itempb.ItemUpdateMsg) (*itempb.ItemMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (srv *ItemServer) Delete(ctx context.Context, req *itempb.ItemDeleteMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (srv *ItemServer) Get(ctx context.Context, req *itempb.ItemGetMsg) (*itempb.ItemGetResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (srv *ItemServer) List(ctx context.Context, req *itempb.ItemListMsg) (*itempb.ItemListResultMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
