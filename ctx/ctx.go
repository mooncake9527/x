package ctxUtil

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ReqId  = "ReqId"
	UserId = "UserId"
	AppKey = "AppKey"
)

type Context struct {
	ginCtx *gin.Context
	context.Context
}

func WithCtx(c *gin.Context) *Context {
	return &Context{
		Context: c.Request.Context(),
		ginCtx:  c,
	}
}

func (x *Context) GetReqId() string {
	return GetReqId(x.ginCtx)
}

func (x *Context) GetUserId() string {
	return GetUserId(x.ginCtx)
}

func (x *Context) GetGinCtx() *gin.Context {
	return x.ginCtx
}

func (x *Context) GetCtx() *context.Context {
	return &x.Context
}

// GetUserId 获取用户的uuid
func GetUserId(c *gin.Context) string {
	return getFromCtx[string](c, UserId)
}

func GetAppKey(c *gin.Context) string {
	return getFromCtx[string](c, AppKey)
}

func SetAppKey(c *gin.Context, appKey string) {
	c.Set(AppKey, appKey)
}

func GetReqId(c *gin.Context) string {
	reqId := c.GetString(ReqId)
	if reqId == "" {
		reqId = uuid.NewString()
		c.Set(reqId, reqId)
	}
	return reqId
}

func getFromCtx[T any](c *gin.Context, key string) T {
	var dft T // default value
	val, exists := c.Get(key)
	if !exists {
		return dft
	}
	if v, ok := val.(T); ok {
		return v
	}
	return dft
}

func tryGetFromCtx[T any](c *gin.Context, keys []string) T {
	var ret T
	for _, key := range keys {
		val, exists := c.Get(key)
		if exists {
			return val.(T)
		}
	}
	return ret
}
