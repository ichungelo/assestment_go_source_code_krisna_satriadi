package fiberhandler

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberInvoice interface {
	CreateInvoice() fiber.Handler
	GetListInvoice() fiber.Handler
	GetInvoiceById() fiber.Handler
	UpdateInvoiceById() fiber.Handler
	DeleteInvoiceById() fiber.Handler
}

type handlerInvoice struct {
	ports.ServiceInvoice
}

func NewInvoiceHandler(sInvoice ports.ServiceInvoice) *handlerInvoice {
	return &handlerInvoice{
		ServiceInvoice: sInvoice,
	}
}

func (h *handlerInvoice) CreateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestCreateInvoice{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceInvoice.CreateInvoice(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerInvoice) GetListInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestGetListInvoice{}

		err := c.QueryParser(&req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		bodyJson, _ := json.MarshalIndent(req, "", "	")
		fmt.Println(string(bodyJson))

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		data, errCode := h.ServiceInvoice.GetListInvoice(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerInvoice) GetInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestGetInvoiceById{}

		invoiceId, err := c.ParamsInt("invoiceId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.InvoiceId = invoiceId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		data, errCode := h.ServiceInvoice.GetInvoiceById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, data, nil, nil)
	}
}

func (h *handlerInvoice) UpdateInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestUpdateInvoiceById{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_FAILED_UNMARSHAL_JSON,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		invoiceId, err := c.ParamsInt("invoiceId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.InvoiceId = invoiceId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceInvoice.UpdateInvoiceById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}

func (h *handlerInvoice) DeleteInvoiceById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := model.RequestDeleteInvoiceById{}

		invoiceId, err := c.ParamsInt("invoiceId", 0)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_PARSE_DATA,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		req.InvoiceId = invoiceId

		err = utils.Validate(req)
		if err != nil {
			utils.Error(err, nil)
			errCode := utils.ErrorCode{
				Code: utils.ERR_VALIDATE,
				Err:  err,
			}

			return fiberpresenter.Presenter(c, nil, nil, &errCode)
		}

		errCode := h.ServiceInvoice.DeleteInvoiceById(&req)
		if errCode != nil {
			return fiberpresenter.Presenter(c, nil, nil, errCode)
		}

		return fiberpresenter.Presenter(c, nil, nil, nil)
	}
}
