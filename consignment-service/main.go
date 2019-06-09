package main

import (
	"context"
	pb "github.com/anymost/micro/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"log"
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

func (service *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	log.Println(req)
	val, err := service.Repository.CreateConsignment(req)
	if err != nil {
		return err
	}
	res = &pb.Response{
		Created:      true,
		Consignment:  val,
	}
	log.Println(res)
	return nil
}

func (service *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res = &pb.Response{
		Consignments: service.Repository.GetAll(),
	}
	return nil
}

func main() {
	server := micro.NewService(micro.Name("go.micro.srv.consignment"), micro.Version("latest"))
	server.Init()
	repository := &Repository{Consignments: make([]*pb.Consignment, 0)}
	service := &Service{Repository: repository}
	pb.RegisterShippingServiceHandler(server.Server(), service)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
