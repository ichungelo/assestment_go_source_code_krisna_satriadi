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
	SUCCESS                         Code = "000"
	ERR_VALIDATE_STRUCT             Code = "001"
	ERR_HASH_PASSWORD               Code = "002"
	ERR_CHECK_PASSWORD              Code = "003"
	ERR_MATCH_CONFIRM_PASSWORD      Code = "004"
	ERR_USER_ALREADY_EXIST          Code = "005"
	ERR_FAILED_CREATE_USER          Code = "006"
	ERR_FAILED_GET_USER             Code = "007"
	ERR_USERNAME_PASSWORD_NOT_MATCH Code = "008"
	ERR_FAILED_UNMARSHAL_JSON       Code = "009"
	ERR_NOT_FOUND                   Code = "099"
	ERR_UNAUTHORIZED                Code = "100"
)

func GetErrorData(errCode ErrorCode) ErrorApi {
	errData := ErrorDataMap[errCode.Code]
	errData.Err = errCode.Err
	return errData
}

var ErrorDataMap = map[Code]ErrorApi{
	"000": {
		Code:           "000",
		Message:        "Success",
		HttpStatusCode: http.StatusOK,
	},
	"001": {
		Code:           "001",
		Message:        "Failed to validate struct",
		HttpStatusCode: http.StatusBadRequest,
	},
	"002": {
		Code:           "002",
		Message:        "Failed to hash password",
		HttpStatusCode: http.StatusInternalServerError,
	},
	"003": {
		Code:           "003",
		Message:        "Failed to check password with confirm password",
		HttpStatusCode: http.StatusBadRequest,
	},
	"004": {
		Code:           "004",
		Message:        "Password different with confirm password",
		HttpStatusCode: http.StatusBadRequest,
	},
	"005": {
		Code:           "005",
		Message:        "User data already exist",
		HttpStatusCode: http.StatusBadRequest,
	},
	"006": {
		Code:           "006",
		Message:        "Failed to create user",
		HttpStatusCode: http.StatusInternalServerError,
	},
	"007": {
		Code:           "007",
		Message:        "Failed to get user",
		HttpStatusCode: http.StatusBadRequest,
	},
	"008": {
		Code:           "008",
		Message:        "Username or password not match",
		HttpStatusCode: http.StatusUnauthorized,
	},
	"009": {
		Code:           "009",
		Message:        "Failed to unmarshal json",
		HttpStatusCode: http.StatusUnauthorized,
	},
	"099": {
		Code:           "099",
		Message:        "Page not found",
		HttpStatusCode: http.StatusNotFound,
	},
	"100": {
		Code:           "100",
		Message:        "You're not login. Please login or sign in first",
		HttpStatusCode: http.StatusUnauthorized,
	},
}
