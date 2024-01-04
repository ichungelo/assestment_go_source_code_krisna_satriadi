package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
)

// Presenter metadata struct
type ResponseMetadata struct {
	Success    *bool               `json:"success"`
	Code       *string             `json:"Code"`
	Error      *string             `json:"error"`
	Message    *string             `json:"message"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

type ResponsePagination struct {
	Total *int `json:"total"`
	Count *int `json:"count"`
	Start *int `json:"start"`
}

// Response Presenter struct
type Response struct {
	Metadata ResponseMetadata `json:"metadata"`
	Result   interface{}      `json:"result"`
}

func Presenter(c *fiber.Ctx, data interface{}, pagination *ResponsePagination, err *utils.ErrorCode) error {
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
		utils.Error(err.Err, nil)

	}

	res := Response{
		Metadata: ResponseMetadata{
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
