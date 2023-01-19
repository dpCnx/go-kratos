package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (r *Response) Error(c *gin.Context, err error) {

	e := errors.FromError(err)
	res := &Response{
		Code: e.Code,
		Msg:  fmt.Sprintf("%s,%s", e.Reason, e.Message),
		Data: nil,
	}
	c.JSON(http.StatusOK, res)

}

func (r *Response) Success(c *gin.Context, data interface{}) {
	res := &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}
