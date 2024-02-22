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

type RouterFiberCustomer interface {
	CreateCustomer() fiber.Handler
	GetListCustomer() fiber.Handler
	UpdateCustomerById() fiber.Handler
	DeleteCustomerById() fiber.Handler
}

type handlerCustomer struct {
	ports.ServiceCustomer
}

func NewCustomerHandler(sCustomer ports.ServiceCustomer) *handlerCustomer {
	return &handlerCustomer{
		ServiceCustomer: sCustomer,
	}
}

func (h *handlerCustomer) CreateCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateCustomer{}

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

		httpError := h.ServiceCustomer.CreateCustomer(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerCustomer) GetListCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, httpError := h.ServiceCustomer.GetListCustomer()
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerCustomer) UpdateCustomerById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateCustomerById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrFailedUnmarshalJson, err)
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		customerId, err := c.ParamsInt("customerId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrParseData, err)
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		req.CustomerId = customerId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceCustomer.UpdateCustomerById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerCustomer) DeleteCustomerById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		customerId, err := c.ParamsInt("customerId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.HttpError{
				Code: utilerrors.ErrParseData,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &httpError)
		}

		req := model.RequestDeleteCustomerById{
			CustomerId: customerId,
		}

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			httpError := utilerrors.NewHttpError(utilerrors.ErrValidate, err)
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		httpError := h.ServiceCustomer.DeleteCustomerById(&req)
		if httpError != nil {
			return fiberpresenter.Presenter(c, nil, nil, httpError)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
