package http

import (
	"context"
	address_grpc "gitlab.com/ilaryonov/fiascli-clean/domain/address/delivery/grpc/address"
	grpc_service "gitlab.com/ilaryonov/fiascli-clean/domain/address/service/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Handler struct {
	server         *grpc.Server
	addressService *grpc_service.AddressService
}

func NewHandler(a *grpc_service.AddressService) *Handler {
	gserver := grpc.NewServer()
	handler := &Handler{
		addressService: a,
		server:         gserver,
	}
	address_grpc.RegisterAddressHandlerServer(gserver, handler)
	reflection.Register(gserver)
	return handler
}

func (h *Handler) GetByGuid(ctx context.Context, guid *address_grpc.GuidRequest) (*address_grpc.Address, error) {
	add, err := h.addressService.GetByGuid(guid.Guid)
	if err != nil {
		//todo log
	}
	result := address_grpc.Address{
		Aoguid:     add.Aoguid,
		Aolevel:    add.Aolevel,
		Formalname: add.Formalname,
		Parentguid: add.Parentguid,
		Shortname:  add.Shortname,
		Postalcode: add.Postalcode,
	}
	return &result, err
}

func (h *Handler) Serve() error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	h.server.Serve(listener)
	return nil
}
