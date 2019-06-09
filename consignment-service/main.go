package main

import (
	"context"
	pb "github.com/anymost/micro/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Port = ":9090"
)

type Repository struct {
	Consignments []*pb.Consignment
}

func (repository *Repository) CreateConsignment(consignment *pb.Consignment) (*pb.Consignment, error) {
	repository.Consignments = append(repository.Consignments, consignment)
	return consignment, nil
}

func (repository *Repository) GetAll() []*pb.Consignment {
	return repository.Consignments
}

type Service struct {
	Repository *Repository
}

func (service *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	res, err := service.Repository.CreateConsignment(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Created:      true,
		Consignment:  res,
		Consignments: service.Repository.Consignments,
	}, nil
}

func (service *Service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{
		Created:      true,
		Consignment:  nil,
		Consignments: service.Repository.Consignments,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatal(err)
	}
	repository := &Repository{Consignments: make([]*pb.Consignment, 0)}
	service := &Service{Repository: repository}
	server := grpc.NewServer()
	pb.RegisterShippingServiceServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("service Running in %s", Port)
	}
}
