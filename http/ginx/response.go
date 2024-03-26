package ginx

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DefaultResp() *Response {
	resp := Response{}
	return &resp
}

func Success(ctx *gin.Context, data interface{}) {
	r := DefaultResp()
	r.Data = data
	ctx.JSON(200, r)
}

func Fail(ctx *gin.Context, err error) {
	resp := DefaultResp()
	if er, ok := err.(*Error); ok {
		resp.Code = er.code
		resp.Msg = er.msg

		ctx.JSON(200, resp)
		return
	}

	ctx.JSON(200, InternalServerError)
}

func FailWithMsg(ctx *gin.Context, msg string) {
	Fail(ctx, NewError(-1, msg))
}
