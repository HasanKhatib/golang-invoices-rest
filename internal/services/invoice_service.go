package services

import (
	"hasankhatib/golang-invoices-rest/internal/models"
	"hasankhatib/golang-invoices-rest/internal/repositories"
)

// InvoiceService represents the service for managing invoices
type InvoiceService struct {
	invoiceRepository *repositories.InvoiceRepository
}

// NewInvoiceService creates a new instance of InvoiceService
func NewInvoiceService(invoiceRepository *repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{invoiceRepository: invoiceRepository}
}

// CreateInvoice creates a new invoice
func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {
	return s.invoiceRepository.CreateInvoice(invoice)
}

// GetInvoiceByID retrieves an invoice by ID
func (s *InvoiceService) GetInvoiceByID(id string) (*models.Invoice, error) {
	return s.invoiceRepository.GetInvoiceByID(id)
}

func (s *InvoiceService) GetAllInvoices() ([]models.Invoice, error) {
	invoices, err := s.invoiceRepository.GetAllInvoices()
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
