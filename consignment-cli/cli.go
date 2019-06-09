package main

import (
	"context"
	"encoding/json"
	pb "github.com/anymost/micro/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"time"
)

const (
	Address = "localhost:9090"
)

func readConsignment() (*pb.Consignment, error) {
	buf, err := ioutil.ReadFile("./consignment.json")
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(buf, &consignment)
	if err != nil {
		return nil, err
	}
	return consignment, nil
}

func createConsignment(client pb.ShippingServiceClient, ctx context.Context) {
	consignment, err := readConsignment()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Consignment)
}

func getConsignments(client pb.ShippingServiceClient, ctx context.Context) {
	res, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatal()
	}
	for _, consignment := range res.Consignments {
		log.Println(consignment)
	}
}

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewShippingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	// createConsignment(client, ctx, consignment)
	getConsignments(client, ctx)
}
