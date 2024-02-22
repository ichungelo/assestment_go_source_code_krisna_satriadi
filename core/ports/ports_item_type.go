package ports

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

type RepositoryItemType interface {
	CreateItemType(itemType *model.ItemType) error
	GetListItemType() ([]model.ResponseGetListItemType, error)
	UpdateItemTypeById(itemType *model.ItemType) error
	DeleteItemTypeById(itemTypeId *int) error
}

type ServiceItemType interface {
	CreateItemType(req *model.RequestCreateItemType) *utilerrors.HttpError
	GetListItemType() ([]model.ResponseGetListItemType, *utilerrors.HttpError)
	UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utilerrors.HttpError
	DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utilerrors.HttpError
}
