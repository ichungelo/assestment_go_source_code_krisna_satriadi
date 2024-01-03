package fiberhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/ports"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

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
		errApi := utils.GetErrorData(errCode)
		utils.Error(errApi.Err, nil)

		var (
			success = false
			errMsg  = errApi.Err.Error()
		)
	
		res := model.Response{
			Metadata: model.ResponseMetadata{
				Success:    &success,
				Code:       &errApi.Code,
				Error:      &errMsg,
				Message:    &errApi.Message,
				Pagination: nil,
			},
			Result: nil,
		}
			// Return status 401 and failed authentication error.
		return c.Status(errApi.HttpStatusCode).JSON(res)
	}
}
