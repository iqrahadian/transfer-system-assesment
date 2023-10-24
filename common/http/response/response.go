package response

import (
	"encoding/json"
	"net/http"

	"github.com/iqrahadian/paperid-assesment/common"
	errMessage "github.com/iqrahadian/paperid-assesment/common/message"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
}

type ErrorResponse struct {
	Cause   string `json:"cause,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func Success(res http.ResponseWriter, message string, httpCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(httpCode)

	if message == "" {
		message = "Success"
	}

	response := SuccessResponse{
		Message: message,
		Data:    data,
		Code:    httpCode,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		// logger.Error("Failed on Encoding the Response", err)

		Error(res, common.Error{err, http.StatusInternalServerError, errMessage.INTERNAL_SERVER_ERROR})

	}
}

func Error(res http.ResponseWriter, errorData common.Error) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(errorData.Code)

	response := ErrorResponse{
		Cause:   errorData.Error.Error(),
		Message: errorData.Message,
		Code:    errorData.Code,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		// logger.Error("Failed on Encoding the Response", err)

		Error(res, common.Error{err, http.StatusInternalServerError, errMessage.INTERNAL_SERVER_ERROR})
	}
}
