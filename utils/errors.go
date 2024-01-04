package utils

import "net/http"

type ErrorApi struct {
	Code           string `json:"code,omitempty"`
	Err            error  `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	HttpStatusCode int    `json:"httpStatusCode,omitempty"`
}

type Code string
type ErrorCode struct {
	Code
	Err error
}

const (
	SUCCESS                   Code = "000"
	ERR_VALIDATE_STRUCT       Code = "001"
	ERR_PARSE_DATA            Code = "002"
	ERR_FAILED_UNMARSHAL_JSON Code = "003"
	ERR_NOT_FOUND             Code = "099"
	ERR_INTERNAL_SERVER_ERROR Code = "100"
)

func GetErrorData(errCode ErrorCode) ErrorApi {
	errData := ErrorDataMap[errCode.Code]
	errData.Err = errCode.Err
	return errData
}

var ErrorDataMap = map[Code]ErrorApi{
	SUCCESS: {
		Code:           "000",
		Message:        "Success",
		HttpStatusCode: http.StatusOK,
	},
	ERR_VALIDATE_STRUCT: {
		Code:           "001",
		Message:        "Failed to validate struct",
		HttpStatusCode: http.StatusBadRequest,
	},
	ERR_PARSE_DATA: {
		Code:           "002",
		Message:        "Failed to parse data",
		HttpStatusCode: http.StatusBadRequest,
	},
	ERR_FAILED_UNMARSHAL_JSON: {
		Code:           "003",
		Message:        "Failed to unmarshal json",
		HttpStatusCode: http.StatusUnauthorized,
	},
	ERR_NOT_FOUND: {
		Code:           "099",
		Message:        "Page not found",
		HttpStatusCode: http.StatusNotFound,
	},
	ERR_INTERNAL_SERVER_ERROR: {
		Code:           "100",
		Message:        "internal server error",
		HttpStatusCode: http.StatusInternalServerError,
	},
}
