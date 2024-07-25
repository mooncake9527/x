package xerror

import "github.com/mooncake9527/x/xerrors/xcode"

// Code 获取错误码。
// 如果没有错误代码，则返回 `xcode.CodeNil`。
func (err *Error) Code() xcode.Code {
	if err == nil {
		return xcode.CodeNil
	}
	if err.code == xcode.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode 使用指定 `code` 更新内部 `code` 。
func (err *Error) SetCode(code xcode.Code) {
	if err == nil {
		return
	}
	err.code = code
}
