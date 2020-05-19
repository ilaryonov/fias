package http

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	address_grpc "github.com/ilaryonov/fias/domain/address/delivery/grpc/address"
	grpc_service "github.com/ilaryonov/fias/domain/address/service/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Handler struct {
	Server         *grpc.Server
	addressService *grpc_service.AddressService
}

func (h *Handler) GetCitiesByTerm(request *address_grpc.TermRequest, stream address_grpc.AddressHandler_GetCitiesByTermServer) error {
	if request.Count == 0 {
		request.Count = 10
	}
	cities, err := h.addressService.GetCitiesByTerm(request.Term, request.Count)
	if err != nil {
		//todo log
	}
	for _, city := range cities {
		stream.Send(&address_grpc.Address{
			Aoguid:     city.Aoguid,
			Aolevel:    city.Aolevel,
			Formalname: city.Formalname,
			Parentguid: city.Parentguid,
			Shortname:  city.Shortname,
			Postalcode: city.Postalcode,
		})
	}
	return nil
}

func (h *Handler) GetAllCities(empty *empty.Empty, stream address_grpc.AddressHandler_GetAllCitiesServer) error {
	cities, err := h.addressService.GetCities()
	if err != nil {
		//todo log
	}
	for _, city := range cities {
		stream.Send(&address_grpc.Address{
			Aoguid:     city.Aoguid,
			Aolevel:    city.Aolevel,
			Formalname: city.Formalname,
			Parentguid: city.Parentguid,
			Shortname:  city.Shortname,
			Postalcode: city.Postalcode,
		})
	}
	return nil
}

func NewHandler(a *grpc_service.AddressService) *Handler {
	/*c, err := credentials.NewServerTLSFromFile("config/Server.pem", "config/Server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}*/
	/*grpc.Creds(c)*/
	gserver := grpc.NewServer()
	handler := &Handler{
		addressService: a,
		Server:         gserver,
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

	h.Server.Serve(listener)
	return nil
}
