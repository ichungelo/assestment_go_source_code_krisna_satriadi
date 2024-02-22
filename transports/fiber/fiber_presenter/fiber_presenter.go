package fiberpresenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

func Presenter(c *fiber.Ctx, data interface{}, pagination *model.ResponsePagination, err *utilerrors.ErrorCode) error {
	var (
		success bool = true
		meta    utilerrors.ErrorApi
		errMsg  *string
	)

	resMeta := utilerrors.ErrorCode{
		Code: utilerrors.Success,
		Err:  nil,
	}
	meta = utilerrors.GetErrorData(resMeta)

	if err != nil {
		meta = utilerrors.GetErrorData(*err)
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
