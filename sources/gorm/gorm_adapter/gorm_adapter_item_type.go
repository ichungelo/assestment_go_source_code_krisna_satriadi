package gormadapter

import (
	"time"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"gorm.io/gorm/clause"
)

func (g *gormAdapter) CreateItemType(itemType *model.ItemType) error {
	result := g.Create(itemType)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gormAdapter) GetListItemType() ([]model.ItemType, error) {
	var (
		data []model.ItemType
	)

	err := g.Model(&model.ItemType{}).Where("is_delete = ?", false).Order("created_at DESC").Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g *gormAdapter) UpdateItemTypeById(itemType *model.ItemType) error {
	err := g.Model(&itemType).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(itemType).Error
	if err != nil {
		return err
	}

	return nil}

func (g *gormAdapter) DeleteItemTypeById(itemTypeId *int) error {
	var (
		isDelete         = true
		deletedTimestamp = time.Now()
	)

	err := g.Model(&model.ItemType{Id: itemTypeId}).Where("is_delete = ?", false).Clauses(clause.Locking{Strength: "UPDATE"}).Updates(model.ItemType{IsDelete: &isDelete, DeletedAt: &deletedTimestamp}).Error
	if err != nil {
		return err
	}

	return nil
}

