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

type RouterFiberItem interface {
	CreateItem() fiber.Handler
	GetListItem() fiber.Handler
	UpdateItemById() fiber.Handler
	DeleteItemById() fiber.Handler
}

type handlerItem struct {
	ports.ServiceItem
}

func NewItemHandler(sItem ports.ServiceItem) *handlerItem {
	return &handlerItem{
		ServiceItem: sItem,
	}
}

func (h *handlerItem) CreateItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateItem{}

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

		httpError := h.ServiceItem.CreateItem(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItem) GetListItem() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, httpError := h.ServiceItem.GetListItem()
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerItem) UpdateItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateItemById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUnmarshalJson, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.ItemId = itemId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceItem.UpdateItemById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerItem) DeleteItemById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteItemById{}

		itemId, err := c.ParamsInt("itemId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.ItemId = itemId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)

			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceItem.DeleteItemById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
