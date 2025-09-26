package commonutils

import "github.com/gin-gonic/gin"

type Response struct {
	Data          interface{}     `json:"data"`
	ErrorResponse []ErrorResponse `json:"error"`
}
type ErrorResponse struct {
	ErrorMsg  string   `json:"error_msg"`
	Field     string   `json:"field"`
	ErrorCode int      `json:"error_code"`
	Vals      []string `json:"vals"`
}

func SendResponse(ctx *gin.Context, statusCode int, data interface{}, errorResponse []ErrorResponse) {
	ctx.JSON(statusCode, Response{
		Data:          data,
		ErrorResponse: errorResponse,
	})
}
func BuildErrorResponse(errorMsg string, errCode int, field string, vals ...string) ErrorResponse {
	return ErrorResponse{
		ErrorMsg:  errorMsg,
		ErrorCode: errCode,
		Field:     field,
		Vals:      vals,
	}
}
