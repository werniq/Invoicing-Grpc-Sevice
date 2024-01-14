package main

import (
	invoice_service "github.com/Chained/invoice-service/github.com/Chained/invoice-service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	InvoiceHttpServerPort int
	InvoiceGrpcServerPort int
	InvoiceGrpcClientPort int

	SMTP struct {
		host string
		port int
		user string
		pwd  string
	}

	client string
}

type Application struct {
	config      config
	infoLogger  *log.Logger
	errorLogger *log.Logger

	invoiceServiceClient invoice_service.InvoiceServiceClient
}

var grpcServer *InvoiceServiceServer

func init() {
	// TODO: load .env variables and configure config
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: %v\n", err)
	}

	log.Println(".env file loaded.")
}

func main() {
	cfg := config{
		SMTP: struct {
			host string
			port int
			user string
			pwd  string
		}{
			// 25, 465, 587, 2525
			host: os.Getenv("YOUR_SMTP_HOST"),
			port: os.Getenv("YOUR_SMTP_PORT"),
			user: os.Getenv("YOUR_SMTP_USERNAME"),
			pwd:  os.Getenv("YOUR_SMTP_USER_PW"),
		},
		InvoiceHttpServerPort: 8082,
		InvoiceGrpcServerPort: 50053,
		client:                "",
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		config:      cfg,
		infoLogger:  infoLog,
		errorLogger: errorLog,
	}

	go app.startGRPCServer()

	if err := app.Serve(); err != nil {
		app.errorLogger.Println(err)
	}
}
