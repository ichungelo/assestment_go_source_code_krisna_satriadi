package gormadapter

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"gorm.io/gorm/clause"
)

func (g *gormAdapter) CreateItem(item *model.Item) error {
	result := g.Create(item)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) GetListItem() ([]model.Item, error) {
	var (
		data []model.Item
	)

	err := g.Model(&model.Item{}).Where("is_delete = ?", false).Order("created_at DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g *gormAdapter) UpdateItemById(item *model.Item) error {
	err := g.Model(&item).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(item).Error
	if err != nil {
		return err
	}

	return nil}

func (g *gormAdapter) DeleteItemById(itemId *int) error {
	var (
		isDelete         = true
		deletedTimestamp = time.Now()
	)

	err := g.Model(&model.Item{Id: itemId}).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(model.Item{IsDelete: &isDelete, DeletedAt: &deletedTimestamp}).Error
	if err != nil {
		return err
	}

	return nil
}

