package gormadapter

import (
	"errors"
	"fmt"
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g *gormAdapter) CreateInvoice(req *model.RequestCreateInvoice) error {
	err := g.Transaction(func(tx *gorm.DB) error {
		invoice := model.Invoice{
			DueDate:    &req.DueDate,
			Subject:    &req.Subject,
			CustomerId: &req.CustomerId,
			Status:     &req.Status,
		}

		err := tx.Clauses(clause.Returning{}).Create(&invoice).Error
		if err != nil {
			return err
		}

		if invoice.Id == nil {
			return err
		}

		items := []model.Quantity{}

		for _, v := range req.Items {
			var (
				itemId = v.ItemId
				count  = v.Count
			)

			item := model.Quantity{
				ItemId:    &itemId,
				InvoiceId: invoice.Id,
				Count:     &count,
			}

			items = append(items, item)
		}

		err = tx.Create(&items).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) GetListInvoice(isDelete bool, limit int, offset int, issueDate *time.Time, subject *string, totalItems *int, customer *string, dueDate *time.Time, InvoiceId *int) (*model.ResponseGetListInvoice, error) {
	var (
		data      []model.ResponseGetListInvoiceResult
		totalData int64
		start     int
		count     int
		total     int
	)

	db := g.Table("invoices").Select("invoices.id AS invoice_id", "invoices.created_at AS issue_date", "invoices.due_date", "invoices.subject", "customers.name AS customer_name", "customers.id AS customer_id", "count(items.id) AS total_items", "invoices.status").
		Joins("LEFT JOIN customers ON customers.id = invoices.customer_id").
		Joins("LEFT JOIN quantities ON quantities.invoice_id = invoices.id").
		Joins("LEFT JOIN items ON quantities.item_id = items.id").
		Order("invoices.created_at DESC").
		Where("invoices.is_delete = ?", isDelete)

	if issueDate != nil {
		startDate := issueDate.Format("2006-01-02")
		endDate := issueDate.Add(24 * time.Hour).Format("2006-01-02")

		db = db.Where("invoices.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	if dueDate != nil {
		startDate := dueDate.Format("2006-01-02")
		endDate := dueDate.Add(24 * time.Hour).Format("2006-01-02")

		db = db.Where("invoices.due_date BETWEEN ? AND ?", startDate, endDate)
	}

	if subject != nil && *subject != "" {
		db = db.Where("invoices.subject LIKE ?", fmt.Sprintf("%%%s%%", *subject))
	}

	if totalItems != nil {
		db = db.Having(`COUNT(items.id) = ?`, totalItems)
	}

	if customer != nil && *customer != "" {
		db = db.Where("customers.name LIKE ?", fmt.Sprintf("%%%s%%", *customer))
	}

	if InvoiceId != nil {
		db = db.Where("invoices.id = ?", InvoiceId)
	}

	if offset != -1 {
		start = offset
	}

	err := db.Group("invoices.id").Count(&totalData).Error
	if err != nil {
		return nil, err
	}

	err = db.Group("invoices.id").Limit(limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, err
	}

	total = int(totalData)
	count = len(data)

	listInvoice := &model.ResponseGetListInvoice{
		Total:  total,
		Count:  count,
		Start:  start,
		Result: data,
	}

	return listInvoice, nil
}

func (g *gormAdapter) GetInvoiceById(invoiceId *int) (*model.ResponseInvoiceById, error) {
	var (
		data = model.ResponseInvoiceById{}
	)

	err := g.Transaction(func(tx *gorm.DB) error {
		err := g.Raw(
			`SELECT invoices.id AS invoice_id, invoices.created_at AS issue_date, invoices.due_date, invoices.subject, customers.name AS customer_name, customers.id AS customer_id, count(items.id) AS total_items, invoices.status
			FROM invoices 
			LEFT JOIN customers ON customers.id = invoices.customer_id
			LEFT JOIN quantities ON quantities.invoice_id = invoices.id
			LEFT JOIN items ON quantities.item_id = items.id
			WHERE invoices.id = ?
			AND invoices.is_delete = false
			GROUP BY invoices.id`, invoiceId).Scan(&data).Error
		if err != nil {
			return err
		}

		err = g.Raw(
			`
			SELECT items.id, items.name, items.unit_price, quantities.count AS quantity, (items.unit_price * quantities.count) as total_price FROM items
			LEFT JOIN quantities ON items.id = quantities.item_id
			LEFT JOIN invoices ON invoices.id = quantities.invoice_id
			WHERE invoices.id = ?;
			`, invoiceId).Scan(&data.Items).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (g *gormAdapter) UpdateInvoiceById(req *model.RequestUpdateInvoiceById) error {
	err := g.Transaction(func(tx *gorm.DB) error {
		invoice := model.Invoice{
			Id:         &req.InvoiceId,
			DueDate:    &req.DueDate,
			Subject:    &req.Subject,
			CustomerId: &req.CustomerId,
			Status:     &req.Status,
		}

		err := tx.Model(&invoice).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(invoice).Error
		if err != nil {
			return err
		}

		updateItems := []model.Quantity{}
		deleteItems := []model.Quantity{}
		for _, v := range req.Items {
			var (
				itemId    = v.ItemId
				count     = v.Count
				invoiceId = req.InvoiceId
			)

			item := model.Quantity{
				ItemId:    &itemId,
				InvoiceId: &invoiceId,
				Count:     count,
			}

			if v.Action == "EDIT" || v.Action == "ADD" {
				if count == nil {
					return fmt.Errorf("count is null")
				}
				updateItems = append(updateItems, item)
			}
			if v.Action == "DELETE" {
				deleteItems = append(deleteItems, item)
			}
		}

		if len(updateItems) > 0 {
			for _, v := range updateItems {
				if err := tx.Model(&v).Where("item_id = ?", v.ItemId).Where("invoice_id = ?", v.InvoiceId).Update("count", v.Count).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						tx.Create(&v)
					}
				}
			}
		}

		if len(deleteItems) > 0 {
			for _, v := range deleteItems {				
				err := tx.Model(&v).Where("item_id = ?", v.ItemId).Where("invoice_id = ?", v.InvoiceId).Delete(&v).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
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
