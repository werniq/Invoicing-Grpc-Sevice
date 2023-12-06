package main

import (
	"fmt"
	pb "github.com/Chained/invoice-service/github.com/Chained/invoice-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var invoiceServiceClient pb.InvoiceServiceClient

// initGrpcClient configures Invoice Microservice Grpc Client
func (app *Application) initGrpcClient() error {
	time.Sleep(time.Second * 5)

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", app.config.InvoiceGrpcClientPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	app.invoiceServiceClient = pb.NewInvoiceServiceClient(conn)

	return nil
}
