package fiberhandler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
	utilvalidator "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_validator"
)

type RouterFiberItemType interface {
	CreateItemType() fiber.Handler
	GetListItemType() fiber.Handler
	UpdateItemTypeById() fiber.Handler
	DeleteItemTypeById() fiber.Handler
}

type handlerItemType struct {
	ports.ServiceItemType
}

func NewItemTypeHandler(sItemType ports.ServiceItemType) *handlerItemType {
	return &handlerItemType{
		ServiceItemType: sItemType,
	}
}

func (h *handlerItemType) CreateItemType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateItemType{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUnmarshalJson, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceItemType.CreateItemType(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItemType) GetListItemType() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, httpError := h.ServiceItemType.GetListItemType()
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerItemType) UpdateItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateItemTypeById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUnmarshalJson, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		itemTypeId, err := c.ParamsInt("itemTypeId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.ItemTypeId = itemTypeId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceItemType.UpdateItemTypeById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItemType) DeleteItemTypeById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteItemTypeById{}

		itemTypeId, err := c.ParamsInt("itemTypeId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.ItemTypeId = itemTypeId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceItemType.DeleteItemTypeById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
