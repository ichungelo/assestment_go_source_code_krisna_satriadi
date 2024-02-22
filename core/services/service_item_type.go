package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

type serviceItemType struct {
	RepositoryItemType ports.RepositoryItemType
}

func NewServiceItemType(rItemType ports.RepositoryItemType) *serviceItemType {
	return &serviceItemType{
		RepositoryItemType: rItemType,
	}
}

func (s *serviceItemType) CreateItemType(req *model.RequestCreateItemType) *utilerrors.ErrorCode {
	var (
		ItemType = model.ItemType{
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.CreateItemType(&ItemType)
	if err != nil {
		utillogger.Error(err, nil)
		errData := utilerrors.ErrorCode{
			Code: utilerrors.ErrFailedCreateItemType,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceItemType) GetListItemType() ([]model.ResponseGetListItemType, *utilerrors.ErrorCode) {
	res, err := s.RepositoryItemType.GetListItemType()
	if err != nil {
		utillogger.Error(err, nil)
		errData := utilerrors.ErrorCode{
			Code: utilerrors.ErrFailedGetItemType,
			Err:  err,
		}
		return nil, &errData
	}

	return res, nil
}

func (s *serviceItemType) UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utilerrors.ErrorCode {
	var (
		itemType = model.ItemType{
			Id:   &req.ItemTypeId,
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.UpdateItemTypeById(&itemType)
	if err != nil {
		utillogger.Error(err, nil)
		errData := utilerrors.ErrorCode{
			Code: utilerrors.ErrFailedUpdateItemType,
			Err:  err,
		}
		return &errData

	}

	return nil
}

func (s *serviceItemType) DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utilerrors.ErrorCode {
	err := s.RepositoryItemType.DeleteItemTypeById(&req.ItemTypeId)
	if err != nil {
		utillogger.Error(err, nil)
		errData := utilerrors.ErrorCode{
			Code: utilerrors.ErrFailedDeleteItemType,
			Err:  err,
		}
		return &errData

	}

	return nil
}
