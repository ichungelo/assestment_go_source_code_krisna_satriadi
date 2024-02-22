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

func (s *serviceItemType) CreateItemType(req *model.RequestCreateItemType) *utilerrors.HttpError {
	var (
		ItemType = model.ItemType{
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.CreateItemType(&ItemType)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedCreateItemType, err)

		return httpError

	}

	return nil
}

func (s *serviceItemType) GetListItemType() ([]model.ResponseGetListItemType, *utilerrors.HttpError) {
	res, err := s.RepositoryItemType.GetListItemType()
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedGetItemType, err)

		return nil, httpError
	}

	return res, nil
}

func (s *serviceItemType) UpdateItemTypeById(req *model.RequestUpdateItemTypeById) *utilerrors.HttpError {
	var (
		itemType = model.ItemType{
			Id:   &req.ItemTypeId,
			Name: &req.Name,
		}
	)

	err := s.RepositoryItemType.UpdateItemTypeById(&itemType)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUpdateItemType, err)

		return httpError

	}

	return nil
}

func (s *serviceItemType) DeleteItemTypeById(req *model.RequestDeleteItemTypeById) *utilerrors.HttpError {
	err := s.RepositoryItemType.DeleteItemTypeById(&req.ItemTypeId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedDeleteItemType, err)

		return httpError

	}

	return nil
}
