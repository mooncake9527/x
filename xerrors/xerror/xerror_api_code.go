package xerror

import (
	"fmt"
	"github.com/mooncake9527/x/xerrors/xcode"
	"strings"
)

// NewCode 用于创建一个自定义错误信息的 error 对象，并包含堆栈信息，并增加错误码对象的输入。
func NewCode(code xcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, separatorSpace),
		code:  code,
	}
}

// NewCodef 用于创建一个自定义错误信息的 error 对象，并包含堆栈信息，并增加错误码对象的输入。
func NewCodef(code xcode.Code, format string, args ...any) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// NewCodeSkip 用于创建一个自定义错误信息的 error 对象，并包含堆栈信息，并增加错误码对象的输入。并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func NewCodeSkip(code xcode.Code, skip int, text ...string) error {
	return &Error{
		stack: callers(skip),
		text:  strings.Join(text, separatorSpace),
		code:  code,
	}
}

// NewCodeSkipf 用于创建一个自定义错误信息的 error 对象，并包含堆栈信息，并增加错误码对象的输入。并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func NewCodeSkipf(code xcode.Code, skip int, format string, args ...any) error {
	return &Error{
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCode 用 error 和文本包裹错误。
func WrapCode(code xcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, separatorSpace),
		code:  code,
	}
}

// WrapCodef 用 error 和格式指定符包裹错误。
func WrapCodef(code xcode.Code, err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// WrapCodeSkip 用 error 和文本包裹错误，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func WrapCodeSkip(code xcode.Code, skip int, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  strings.Join(text, separatorSpace),
		code:  code,
	}
}

// WrapCodeSkipf 用 error 和格式指定符包裹错误，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。
func WrapCodeSkipf(code xcode.Code, skip int, err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(skip),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

// Code 获取 error 中的错误码接口。
func Code(err error) xcode.Code {
	if err == nil {
		return xcode.CodeNil
	}
	if e, ok := err.(ICode); ok {
		return e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return Code(e.Unwrap())
	}
	return xcode.CodeNil
}

// HasCode 检查并报告 `err` 在其链接错误中是否具有 `code`。
func HasCode(err error, code xcode.Code) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(ICode); ok {
		return code == e.Code()
	}
	if e, ok := err.(IUnwrap); ok {
		return HasCode(e.Unwrap(), code)
	}
	return false
}
