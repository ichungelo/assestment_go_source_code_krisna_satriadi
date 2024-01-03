package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryItemType interface {
	CreateItemType(itemType *model.ItemType) error
	GetListItemType() ([]model.ItemType, error)
	UpdateItemTypeById(itemType *model.ItemType) error
	DeleteItemTypeById(itemTypeId *int) error
}

type ServiceItemType interface {
	CreateItemType(req *model.RequestCreateItemType) *utils.ErrorCode
	GetListItemType() ([]model.ItemType, *utils.ErrorCode)
	UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utils.ErrorCode
	DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utils.ErrorCode
}