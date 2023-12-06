package main

import (
	"context"
	"fmt"
	pb "github.com/Chained/invoice-service/github.com/Chained/invoice-service"
	"github.com/gin-gonic/gin"
	"github.com/go-pdf/fpdf"
	"github.com/golang/protobuf/jsonpb"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	Product   string    `json:"product"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerateInvoice creates Order as a pdf file, and emails it to the recipient
func (app *Application) GenerateInvoice(ctx *gin.Context) {
	var order *pb.InvoiceRequest

	err := jsonpb.Unmarshal(ctx.Request.Body, order)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "unmarshalling json: " + err.Error(),
		})
		return
	}

	res, err := grpcServer.GenerateInvoice(context.Background(), order)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "generating an invoice: " + err.Error(),
		})
	}

	app.infoLogger.Println(res)

	// if err == nil, then -> res != nil
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "invoice created and successfully sent.",
	})
}

// createInvoice generates PDF version of the invoice
// creates a new PDF document with portrait
// orientation ("P"), millimeter units ("mm"), A4 page size, and an empty title.
//
// sets the font for the title to Arial, bold, and size 16.
//
// print the title: pdf.Cell(40, 10, "Invoice") adds a cell with the text "Invoice" to the PDF.
//
// pdf.Ln(10) adds a line break of 10 units.
//
// pdf.SetFont("Arial", "", 12) sets the font for the rest of the document to Arial and size 12.
//
// Order struct prints the order details and customer information.
//
// save the PDF to a file: pdf.OutputFileAndClose("${orderId}.pdf") saves the PDF to a file named "${orderId}.pdf".
func (app *Application) createInvoice(order *pb.InvoiceRequest) error {
	// Create a new PDF instance
	pdf := fpdf.New("P", "mm", "A4", "16")

	// Add a new page
	pdf.AddPage()

	// Set font for the text
	pdf.SetFont("Arial", "b", 16)

	// Print order details
	pdf.Cell(40, 10, fmt.Sprintf("Order ID: %d", order.Id))
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Product: %s", order.Product))
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Quantity: %d", order.Quantity))
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Amount: $%d", order.Amount))
	pdf.Ln(8)

	// Line break
	pdf.Ln(10)

	// Print customer information
	pdf.Cell(40, 10, "Customer Information:")
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Name: %s %s", order.Firstname, order.Lastname))
	pdf.Ln(8)

	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", order.Email))
	pdf.Ln(8)

	// Line break
	pdf.Ln(10)

	// Print timestamp
	//pdf.Cell(40, 10, fmt.Sprintf("Order Date: %s", order.CreatedAt))
	//pdf.Ln(8)

	// Set font for the timestamp if uncommented
	// pdf.SetFont("Arial", "", 12)

	// Specify the path for the generated PDF
	invoicePath := fmt.Sprintf("./invoices/%d.pdf", order.Id)

	// Output the PDF to the specified file
	return pdf.OutputFileAndClose(invoicePath)
}
