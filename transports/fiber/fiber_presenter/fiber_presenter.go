package fiberpresenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utilerrors "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_errors"
)

func Presenter(c *fiber.Ctx, data interface{}, pagination *model.ResponsePagination, httpError *utilerrors.HttpError) error {
	var (
		success bool = true
		meta    utilerrors.ErrorApi
		errMsg  *string
	)

	resMeta := utilerrors.HttpError{
		Code: utilerrors.Success,
		Err:  nil,
	}
	meta = utilerrors.GetErrorData(resMeta)

	if httpError != nil {
		meta = utilerrors.GetErrorData(*httpError)
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
