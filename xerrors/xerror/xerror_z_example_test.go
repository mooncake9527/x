package xerror_test

import (
	"errors"
	"fmt"
	"github.com/mooncake9527/x/xerrors/xcode"
	"github.com/mooncake9527/x/xerrors/xerror"
)

func ExampleNewCode() {
	err := xerror.NewCode(xcode.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(xerror.Code(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := xerror.NewCodef(xcode.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(xerror.Code(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := xerror.WrapCode(xcode.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(xerror.Code(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := xerror.WrapCodef(xcode.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(xerror.Code(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleEqual() {
	err1 := errors.New("permission denied")
	err2 := xerror.New("permission denied")
	err3 := xerror.NewCode(xcode.CodeNotAuthorized, "permission denied")
	fmt.Println(xerror.Equal(err1, err2))
	fmt.Println(xerror.Equal(err2, err3))

	// Output:
	// true
	// false
}

func ExampleIs() {
	err1 := errors.New("permission denied")
	err2 := xerror.Wrap(err1, "operation failed")
	fmt.Println(xerror.Is(err1, err1))
	fmt.Println(xerror.Is(err2, err2))
	fmt.Println(xerror.Is(err2, err1))
	fmt.Println(xerror.Is(err1, err2))

	// Output:
	// false
	// true
	// true
	// false
}
