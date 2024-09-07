package dcontext

//
//type Context struct {
//	ginCtx *gin.Context
//	context.Context
//}
//
//func WithCtx(c *gin.Context) *Context {
//	return &Context{
//		Context: c.Request.Context(),
//		ginCtx:  c,
//	}
//}
//
//func (x *Context) GetReqId() string {
//	return getReqId(x.ginCtx)
//}
//
//func (x *Context) GetCompanyId() string {
//	return getCompanyId(x.ginCtx)
//}
//
//func (x *Context) GetUserId() string {
//	return getUserId(x.ginCtx)
//}
//
//func (x *Context) GetGinCtx() *gin.Context {
//	return x.ginCtx
//}
//
//func (x *Context) GetCtx() *context.Context {
//	return &x.Context
//}
//
//func getCompanyId(c *gin.Context) string {
//	companyId, exists := c.Get("companyId")
//	if !exists {
//		return ""
//	}
//	if companyIdStr, ok := companyId.(string); ok {
//		return companyIdStr
//	}
//	return ""
//}
//
//func getUserId(c *gin.Context) string {
//	userId, exists := c.Get("userId")
//	if !exists {
//		return ""
//	}
//	if userIdStr, ok := userId.(string); ok {
//		return userIdStr
//	}
//	return ""
//}
//
//func getReqId(c *gin.Context) string {
//	reqId := c.GetString(consts.REQ_ID)
//	if reqId == "" {
//		reqId = uuid.NewString()
//		c.Set(consts.REQ_ID, reqId)
//	}
//	return reqId
//}
