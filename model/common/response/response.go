package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0 // 成功状态码
	ERROR   = 1 // 失败状态码
)

/*
*
统一响应格式方法
* @code: 状态码
* @data: 返回数据
* @msg:  返回信息
* @c:    gin上下文
*/
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

/**
* 成功响应
* @message: 返回信息
* @c:       gin上下文
 */
func SuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

/*
* 成功响应
* @data:    返回数据
* @message: 返回信息
* @c:       gin上下文
 */
func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

/**
* 失败响应
* @message: 返回信息
* @c:       gin上下文
 */
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

/*
* 失败响应
* @data:    返回数据
* @message: 返回信息
* @c:       gin上下文
 */
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
