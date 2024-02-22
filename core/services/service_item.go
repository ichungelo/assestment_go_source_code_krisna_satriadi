package services

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
)

type serviceItem struct {
	RepositoryItem ports.RepositoryItem
}

func NewServiceItem(rItem ports.RepositoryItem) *serviceItem {
	return &serviceItem{
		RepositoryItem: rItem,
	}
}

func (s *serviceItem) CreateItem(req *model.RequestCreateItem) *utilerrors.HttpError {
	var (
		item = model.Item{
			Name:       &req.Name,
			UnitPrice:  &req.UnitPrice,
			ItemTypeId: &req.ItemTypeId,
		}
	)

	err := s.RepositoryItem.CreateItem(&item)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedCreateItem, err)

		return httpError
	}

	return nil
}

func (s *serviceItem) GetListItem() ([]model.ResponseGetListItem, *utilerrors.HttpError) {
	res, err := s.RepositoryItem.GetListItem()
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedGetItem, err)

		return nil, httpError
	}

	return res, nil
}

func (s *serviceItem) UpdateItemById(req *model.RequestUpdateItemById) *utilerrors.HttpError {
	var (
		item = model.Item{
			Id:         &req.ItemId,
			Name:       &req.Name,
			UnitPrice:  &req.UnitPrice,
			ItemTypeId: &req.ItemTypeId,
		}
	)

	err := s.RepositoryItem.UpdateItemById(&item)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUpdateItem, err)

		return httpError

	}

	return nil
}

func (s *serviceItem) DeleteItemById(req *model.RequestDeleteItemById) *utilerrors.HttpError {
	err := s.RepositoryItem.DeleteItemById(&req.ItemId)
	if err != nil {
		utillogger.Error(err, nil)
		httpError := utilerrors.NewHttpError(utilerrors.ErrFailedDeleteItem, err)

		return httpError

	}

	return nil
}
