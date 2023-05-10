package repositories

import (
	"hasankhatib/golang-invoices-rest/internal/models"
	"gorm.io/gorm"
)

// InvoiceRepository represents the repository for managing invoices
type InvoiceRepository struct {
	db *gorm.DB
}

// NewInvoiceRepository creates a new instance of InvoiceRepository
func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// CreateInvoice creates a new invoice
func (r *InvoiceRepository) CreateInvoice(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

// GetInvoiceByID retrieves an invoice by ID
func (r *InvoiceRepository) GetInvoiceByID(id string) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.First(&invoice, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

// GetAllInvoices retrieves all invoices
func (r *InvoiceRepository) GetAllInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
