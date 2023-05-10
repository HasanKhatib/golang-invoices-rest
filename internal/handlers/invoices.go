package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
	"hasankhatib/golang-invoices-rest/internal/models"
	"hasankhatib/golang-invoices-rest/internal/services"
)

// InvoicesHandler handles the requests related to invoices
type InvoicesHandler struct {
	invoiceService services.InvoiceService
}

// NewInvoiceHandler creates a new InvoicesHandler
func NewInvoiceHandler(invoiceService services.InvoiceService) *InvoicesHandler {
	return &InvoicesHandler{
		invoiceService: invoiceService,
	}
}

// GetAllInvoices returns all invoices
func (h *InvoicesHandler) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
	// Fetch the invoices from the invoice service
	invoices, err := h.invoiceService.GetAllInvoices()
	if err != nil {
		log.Printf("failed to get invoices: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the invoices to JSON
	response, err := json.Marshal(invoices)
	if err != nil {
		log.Printf("failed to marshal invoices: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetInvoiceByID returns an invoice by ID
func (h *InvoicesHandler) GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	// Extract the invoice ID from the request URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Fetch the invoice from the invoice service
	invoice, err := h.invoiceService.GetInvoiceByID(id)
	if err != nil {
		log.Printf("failed to get invoice: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if the invoice exists
	if invoice == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Convert the invoice to JSON
	response, err := json.Marshal(invoice)
	if err != nil {
		log.Printf("failed to marshal invoice: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateInvoice creates a new invoice
func (h *InvoicesHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to extract the invoice details
	var invoice models.Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		log.Printf("failed to parse request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the invoice data
	if invoice.ID == "" || invoice.Description == "" || invoice.Amount <= 0 {
		log.Println("invalid invoice data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the invoice using the invoice service
	err = h.invoiceService.CreateInvoice(&invoice)
	if err != nil {
		log.Printf("failed to create invoice: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusCreated)
}
