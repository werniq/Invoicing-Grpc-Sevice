# Invoicing Microservice

The Invoicing Microservice is a gRPC-based service designed to send emails to users based on their orders and generate PDF invoices. It utilizes the `invoice.proto` file for defining the gRPC service.

## Features

- **Email Notifications**: Sends email notifications to users based on their orders.
- **PDF Generation**: Generates PDF invoices and saves them to the `cmd/invoices` directory.
