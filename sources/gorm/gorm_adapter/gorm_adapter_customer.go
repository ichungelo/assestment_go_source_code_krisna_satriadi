package gormadapter

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"gorm.io/gorm/clause"
)

func (g *gormAdapter) CreateCustomer(customer *model.Customer) error {
	result := g.Create(customer)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) GetListCustomer() ([]model.ResponseGetListCustomer, error) {
	var (
		data []model.ResponseGetListCustomer
	)

	err := g.Model(&model.Customer{}).Where("is_delete = ?", false).Order("created_at DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g *gormAdapter) UpdateCustomerById(customer *model.Customer) error {
	err := g.Model(&customer).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) DeleteCustomerById(customerId *int) error {
	var (
		isDelete         = true
		deletedTimestamp = time.Now()
	)

	err := g.Model(&model.Customer{Id: customerId}).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(model.Customer{IsDelete: &isDelete, DeletedAt: &deletedTimestamp}).Error
	if err != nil {
		return err
	}

	return nil
}
