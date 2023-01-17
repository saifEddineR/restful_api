package helper

import "strings"

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"data"`
	Data    interface{} `json:"data"`
}
// Empty objectObj is used when data doesnt want to be null on json
type EmptyObj struct {
}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}
//BuildErrorResponse method is to inject data value to dynamic failted response

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err,"\n")
	res:= Response{
		Status:false,
		Message:message,
		Errors:splittedError,
		Data:data,
	}
	return res
}