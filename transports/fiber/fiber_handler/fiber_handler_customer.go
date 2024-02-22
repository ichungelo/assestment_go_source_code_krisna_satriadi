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
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrFailedUnmarshalJson,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrValidate,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceCustomer.CreateCustomer(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerCustomer) GetListCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {

		data, errCode := h.ServiceCustomer.GetListCustomer()
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
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
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrFailedUnmarshalJson,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		customerId, err := c.ParamsInt("customerId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrParseData,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.CustomerId = customerId

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrValidate,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceCustomer.UpdateCustomerById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerCustomer) DeleteCustomerById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		customerId, err := c.ParamsInt("customerId", 0)
		if err != nil {
			utillogger.Error(err, nil)
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrParseData,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req := model.RequestDeleteCustomerById{
			CustomerId: customerId,
		}

		err = utilvalidator.Validate(req)
		if err != nil {
			utillogger.Error(err, nil)
			errCode := utilerrors.ErrorCode{
				Code: utilerrors.ErrValidate,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceCustomer.DeleteCustomerById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
