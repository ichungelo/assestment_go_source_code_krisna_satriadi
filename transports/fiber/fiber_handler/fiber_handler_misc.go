package fiberhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

type RouterFiberMisc interface {
	NotFound() fiber.Handler
}

type handlerMisc struct {
	ports.ServiceMisc
}

func NewMiscHandler(sMisc ports.ServiceMisc) *handlerMisc {
	return &handlerMisc{
		ServiceMisc: sMisc,
	}
}

func (h *handlerMisc) NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Return HTTP 404 status and JSON response.
		errCode := utils.ErrorCode{
			Code: utils.ERR_NOT_FOUND,
			Err:  errors.New("page not found"),
		}

		return fiberpresenter.Presenter(c, nil, nil, &errCode)
	}
}
