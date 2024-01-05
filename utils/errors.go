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
	ERR_VALIDATE              Code = "001"
	ERR_PARSE_DATA            Code = "002"
	ERR_FAILED_UNMARSHAL_JSON Code = "003"

	ERR_FAILED_CREATE_CUSTOMER Code = "004"
	ERR_FAILED_GET_CUSTOMER    Code = "005"
	ERR_FAILED_UPDATE_CUSTOMER Code = "006"
	ERR_FAILED_DELETE_CUSTOMER Code = "007"

	ERR_FAILED_CREATE_INVOICE Code = "008"
	ERR_FAILED_GET_INVOICE    Code = "009"
	ERR_FAILED_UPDATE_INVOICE Code = "010"
	ERR_FAILED_DELETE_INVOICE Code = "011"

	ERR_FAILED_CREATE_ITEM_TYPE Code = "012"
	ERR_FAILED_GET_ITEM_TYPE    Code = "013"
	ERR_FAILED_UPDATE_ITEM_TYPE Code = "014"
	ERR_FAILED_DELETE_ITEM_TYPE Code = "015"

	ERR_FAILED_CREATE_ITEM Code = "016"
	ERR_FAILED_GET_ITEM    Code = "017"
	ERR_FAILED_UPDATE_ITEM Code = "018"
	ERR_FAILED_DELETE_ITEM Code = "019"

	ERR_FAILED_DELETE_QUANTITY Code = "020"
	ERR_PARSE_DATE Code= "021"

	ERR_NOT_FOUND             Code = "099"
	ERR_INTERNAL_SERVER_ERROR Code = "100"
	ERR_GENERAL               Code = "999"
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
	ERR_VALIDATE: {
		Code:           "001",
		Message:        "Failed to validate",
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

	ERR_FAILED_CREATE_CUSTOMER: {
		Code:           "004",
		Message:        "Failed to create customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_GET_CUSTOMER: {
		Code:           "005",
		Message:        "Failed to get customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_UPDATE_CUSTOMER: {
		Code:           "006",
		Message:        "Failed to update customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_DELETE_CUSTOMER: {
		Code:           "007",
		Message:        "Failed to delete customer",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ERR_FAILED_CREATE_INVOICE: {
		Code:           "008",
		Message:        "Failed to create invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_GET_INVOICE: {
		Code:           "009",
		Message:        "Failed to get invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_UPDATE_INVOICE: {
		Code:           "010",
		Message:        "Failed to update invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_DELETE_INVOICE: {
		Code:           "011",
		Message:        "Failed to delete invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ERR_FAILED_CREATE_ITEM_TYPE: {
		Code:           "012",
		Message:        "Failed to create item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_GET_ITEM_TYPE: {
		Code:           "013",
		Message:        "Failed to get item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_UPDATE_ITEM_TYPE: {
		Code:           "014",
		Message:        "Failed to update item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_DELETE_ITEM_TYPE: {
		Code:           "015",
		Message:        "Failed to delete item type",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ERR_FAILED_CREATE_ITEM: {
		Code:           "016",
		Message:        "Failed to create item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_GET_ITEM: {
		Code:           "017",
		Message:        "Failed to get item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_UPDATE_ITEM: {
		Code:           "018",
		Message:        "Failed to update item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_FAILED_DELETE_ITEM: {
		Code:           "019",
		Message:        "Failed to delete item",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ERR_FAILED_DELETE_QUANTITY: {
		Code:           "020",
		Message:        "Failed to delete quantity",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ERR_PARSE_DATE: {
		Code:           "021",
		Message:        "Failed to parse date",
		HttpStatusCode: http.StatusInternalServerError,
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
	ERR_GENERAL: {
		Code:           "999",
		Message:        "error happen",
		HttpStatusCode: http.StatusInternalServerError,
	},
}
