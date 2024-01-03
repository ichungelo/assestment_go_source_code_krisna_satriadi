package gormadapter

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"gorm.io/gorm/clause"
)

func (g *gormAdapter) CreateInvoice(invoice *model.Invoice) error {
	result := g.Create(invoice)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) GetListInvoice(isDelete bool, limit int, offset int, issueDate *time.Time, subject *string, totalItems *int, customer *string, dueDate *time.Time, InvoiceId *int) (total int, count int, start int, listInvoice []model.Invoice, err error) {
	var (
		totalData int64
	)

	db := g.Model(&model.Invoice{}).
		Where("invoices.is_delete = ?", isDelete).
		Order("invoices.created_at DESC")

	if issueDate != nil {
		db = db.Where("issue_date = ?", *issueDate)
	}

	if dueDate != nil {
		db = db.Where("created_at = ?", *dueDate)
	}

	if subject != nil {
		db = db.Where("subjects LIKE \"%?%\"", subject)
	}

	if totalItems != nil {
		db = db.Preload("Quantities").Where("sum(quantities.count) = ?", totalItems)
	}

	if customer != nil {
		db = db.Joins("INNER JOIN customers ON customer.id = invoices.customer_id").Where("customers.name LIKE \"%?%\"", customer)
	}

	if InvoiceId != nil {
		db = db.Where("invoices.id = ?", InvoiceId)
	}

	if offset != -1 {
		start = offset
	}

	err = db.Count(&totalData).Error
	if err != nil {
		return 0, 0, 0, nil, err
	}

	err = db.Limit(limit).Offset(offset).Find(&listInvoice).Error
	if err != nil {
		return 0, 0, 0, nil, err
	}

	total = int(totalData)
	count = len(listInvoice)
	return
}

func (g *gormAdapter) UpdateInvoiceById(invoice *model.Invoice) error {
	err := g.Model(&invoice).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(invoice).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) DeleteInvoiceById(invoiceId *int) error {
	var (
		isDelete         = true
		deletedTimestamp = time.Now()
	)

	err := g.Model(&model.Invoice{Id: invoiceId}).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(model.Invoice{IsDelete: &isDelete, DeletedAt: &deletedTimestamp}).Error
	if err != nil {
		return err
	}

	return nil
}