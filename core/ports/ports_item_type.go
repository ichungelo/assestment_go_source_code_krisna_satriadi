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
	CreateItemType(req *model.RequestCreateItemType) *utilerrors.ErrorCode
	GetListItemType() ([]model.ResponseGetListItemType, *utilerrors.ErrorCode)
	UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utilerrors.ErrorCode
	DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utilerrors.ErrorCode
}
