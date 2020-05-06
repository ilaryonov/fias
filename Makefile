import:
	go build -o fiascli && ./fiascli checkupdates
grpcGenAddress:
	protoc domain/address/delivery/grpc/address/address.proto --go_out=plugins=grpc:.