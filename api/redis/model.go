package redis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	ErrorMessage string     `json:"error_message,omitempty"`
	ErrorCode    errrorCode `json:"error_code,omitempty"`
}

type errrorCode struct {
	Code int    `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Response struct {
	Response any            `json:"response,omitempty"`
	Message  *string        `json:"message,omitempty"`
	Error    *errorResponse `json:"error,omitempty"`
}

func JsonResponse[T any](w http.ResponseWriter, res T) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		fmt.Println("cannot marhal the response, ", err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("cannot write the json response, ", err)
	}
}

func CreateResponse(response any, message string, err error, code int) Response {
	var e *errorResponse
	if err != nil {
		e = &errorResponse{
			ErrorMessage: err.Error(),
			ErrorCode: errrorCode{
				Code: code,
				Name: http.StatusText(code),
			},
		}
	}

	return Response{
		Response: &response,
		Message:  &message,
		Error:    e,
	}

}

func ErrorResponse(writer http.ResponseWriter, request *http.Request, err error, status int) bool {
	if err != nil {
		writer.WriteHeader(status)
		JsonResponse(writer,
			CreateResponse(
				nil,
				fmt.Sprintf("issue with the request [%s]", err),
				err,
				status,
			))
		return true
	}

	return false
}
