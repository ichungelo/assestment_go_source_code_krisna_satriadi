package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RepositoryItem interface {
	CreateItem(item *model.Item) error
	GetListItem() ([]model.Item, error)
	UpdateItemById(item *model.Item) error
	DeleteItemById(itemId *int) error
}

type ServiceItem interface {
	CreateItem(req *model.RequestCreateItem) *utils.ErrorCode
	GetListItem() ([]model.Item, *utils.ErrorCode)
	UpdateItemById(req *model.RequestUpdateItemById) *utils.ErrorCode
	DeleteItemById(req *model.RequestDeleteItemById) *utils.ErrorCode
}