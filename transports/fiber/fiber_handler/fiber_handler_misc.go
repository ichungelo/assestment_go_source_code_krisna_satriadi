package fiberhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	fiberpresenter "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/transports/fiber/fiber_presenter"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
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
		errCode := utilerrors.ErrorCode{
			Code: utilerrors.ErrNotFound,
			Err:  errors.New("page not found"),
		}

		return fiberpresenter.Presenter(c, nil, nil, &errCode)
	}
}
