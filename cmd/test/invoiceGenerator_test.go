package test

import (
	"context"
	invoice_service "github.com/Chained/invoice-service/github.com/Chained/invoice-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

type TestInvoiceServiceServer struct {
	invoice_service.UnimplementedInvoiceServiceServer
}

//
//const bufSize = 1024 * 1024
//
//var lis *bufconn.Listener
//
//func init() {
//	lis = bufconn.Listen(bufSize)
//
//	s := grpc.NewServer()
//	invoice_service.RegisterInvoiceServiceServer(s, &TestInvoiceServiceServer{})
//	go func() {
//		if err := s.Serve(lis); err != nil {
//			log.Fatalf("Server exited with the error: %v\n", err)
//		}
//	}()
//}
//
//func bufDialer(context.Context, string) (net.Conn, error) {
//	return lis.Dial()
//}

func TestGenerateInvoice(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v\n", err)
	}

	defer conn.Close()

	client := invoice_service.NewInvoiceServiceClient(conn)

	order := &invoice_service.InvoiceRequest{
		Id:        1,
		Quantity:  2,
		Amount:    5,
		Product:   "Finances ERP Services",
		Firstname: "Oleksandr",
		Lastname:  "Matviienko",
		Email:     "qniwwwersss@gmail.com",
	}

	_, err = client.GenerateInvoice(context.Background(), order)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
