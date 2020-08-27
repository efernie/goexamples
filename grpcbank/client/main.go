package main

import (
	"context"
	"log"
	"time"

	"github.com/efernie/goexamples/grpcbank/pkg/proto/credit"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := credit.NewCreditServiceClient(conn)

	request := &credit.CreditRequest{Amount: 1990.01}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Credit(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.GetConfirmation())
}
