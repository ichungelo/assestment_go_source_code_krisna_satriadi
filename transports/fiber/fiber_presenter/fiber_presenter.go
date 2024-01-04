package fiberpresenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

func Presenter(c *fiber.Ctx, data interface{}, pagination *model.ResponsePagination, err *utils.ErrorCode) error {
	var (
		success bool = true
		meta    utils.ErrorApi
		errMsg  *string
	)

	resMeta := utils.ErrorCode{
		Code: utils.SUCCESS,
		Err:  nil,
	}
	meta = utils.GetErrorData(resMeta)

	if err != nil {
		meta = utils.GetErrorData(*err)
		message := meta.Err.Error()
		errMsg = &message
		success = false
	}

	res := model.Response{
		Metadata: model.ResponseMetadata{
			Success:    &success,
			Code:       &meta.Code,
			Error:      errMsg,
			Message:    &meta.Message,
			Pagination: pagination,
		},
		Result: data,
	}

	return c.Status(meta.HttpStatusCode).JSON(res)
}
