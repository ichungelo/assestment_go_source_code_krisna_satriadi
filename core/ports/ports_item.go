package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

type RepositoryItem interface {
	CreateItem(item *model.Item) error
	GetListItem() ([]model.ResponseGetListItem, error)
	UpdateItemById(item *model.Item) error
	DeleteItemById(itemId *int) error
}

type ServiceItem interface {
	CreateItem(req *model.RequestCreateItem) *utilerrors.ErrorCode
	GetListItem() ([]model.ResponseGetListItem, *utilerrors.ErrorCode)
	UpdateItemById(req *model.RequestUpdateItemById) *utilerrors.ErrorCode
	DeleteItemById(req *model.RequestDeleteItemById) *utilerrors.ErrorCode
}
