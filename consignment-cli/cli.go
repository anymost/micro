package main

import (
	"context"
	"encoding/json"
	pb "github.com/anymost/micro/consignment-service/proto/consignment"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
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
	log.Println(client.CreateConsignment(ctx, consignment))
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
	if err := cmd.Init(); err != nil {
		log.Fatal(err)
	}
	c := pb.NewShippingServiceClient("go.micro.srv.consignment", client.NewClient())
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	createConsignment(c, ctx)
	// getConsignments(c, ctx)
}
