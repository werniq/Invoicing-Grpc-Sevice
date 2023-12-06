package main

import (
	"context"
	"fmt"
	pb "github.com/Chained/invoice-service/github.com/Chained/invoice-service"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

type InvoiceServiceServer struct {
	App *Application
	pb.UnimplementedInvoiceServiceServer
}

// startGRPCServer initializes and starts the gRPC server.
// It listens on the specified port for incoming gRPC connections,
// registers the AuthServiceServer implementation, and serves requests.
// The function logs server initialization information and any errors
// encountered during the server's lifecycle.
func (app *Application) startGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", app.config.InvoiceGrpcServerPort))
	if err != nil {
		app.errorLogger.Printf("failed to listen: %v\n", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	pb.RegisterInvoiceServiceServer(s, &InvoiceServiceServer{App: app})

	app.infoLogger.Printf("gRPC server is listening on port %d...\n", app.config.InvoiceGrpcServerPort)

	if err := s.Serve(lis); err != nil {
		app.errorLogger.Printf("failed to serve: %v\n", err)
	}
}

func (s *InvoiceServiceServer) GenerateInvoice(ctx context.Context, invoiceRequest *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	err := s.App.createInvoice(invoiceRequest)
	if err != nil {
		fmt.Printf("error 1")

		return &pb.InvoiceResponse{
			Error:   true,
			Message: "error creating invoice: " + err.Error(),
		}, err
	}

	attachments := []string{
		fmt.Sprintf("./invoices/%d.pdf", invoiceRequest.Id),
	}

	err = s.App.SendMail("no-reply@chained.com",
		invoiceRequest.Email,
		"Invoice for the order: "+strconv.Itoa(int(invoiceRequest.Id)),
		"invoice",
		attachments,
		nil,
	)

	if err != nil {
		fmt.Printf("error 2")

		return &pb.InvoiceResponse{
			Error:   true,
			Message: "error sending email: " + err.Error(),
		}, err
	}

	return &pb.InvoiceResponse{
		Error:   false,
		Message: "Successfully generated and sent an invoice.",
	}, nil
}
