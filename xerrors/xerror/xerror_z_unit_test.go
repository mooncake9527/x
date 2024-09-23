package xerror_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mooncake9527/x/xerrors/xcode"
	"github.com/mooncake9527/x/xerrors/xerror"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	assert.NotNil(t, xerror.New(""))
	assert.Nil(t, xerror.Wrap(nilError(), "test"))
}

func Test_New(t *testing.T) {
	err1 := xerror.New("1")
	assert.NotNil(t, err1)
	assert.Equal(t, err1.Error(), "1")

	err2 := xerror.Newf("%d", 1)
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "1")

	err3 := xerror.NewSkipf(1, "%d", 1)
	assert.NotNil(t, err3)
	assert.Equal(t, err3.Error(), "1")
}

func Test_Wrap(t *testing.T) {
	err1 := errors.New("1")
	err1 = xerror.Wrap(err1, "2")
	err1 = xerror.Wrap(err1, "3")
	assert.NotNil(t, err1)
	assert.Equal(t, err1.Error(), "3: 2: 1")

	err2 := xerror.New("1")
	err2 = xerror.Wrap(err2, "")
	assert.NotNil(t, err2)
	assert.Equal(t, err2.Error(), "1")
}

func Test_Wrapf(t *testing.T) {
	err1 := errors.New("1")
	err1 = xerror.Wrapf(err1, "%d", 2)
	err1 = xerror.Wrapf(err1, "%d", 3)
	assert.NotNil(t, err1)
	assert.Equal(t, err1.Error(), "3: 2: 1")

	err2 := xerror.New("1")
	err2 = xerror.Wrapf(err2, "")
	assert.NotNil(t, err2, nil)
	assert.Equal(t, err2.Error(), "1")
}

func Test_WrapSkip(t *testing.T) {
	err1 := errors.New("1")
	err1 = xerror.WrapSkip(1, err1, "2")
	err1 = xerror.WrapSkip(1, err1, "3")
	assert.NotNil(t, err1, nil)
	assert.Equal(t, err1.Error(), "3: 2: 1")

	err2 := xerror.New("1")
	err2 = xerror.WrapSkip(1, err2, "")
	assert.NotNil(t, err2, nil)
	assert.Equal(t, err2.Error(), "1")
}

func Test_WrapSkipf(t *testing.T) {
	err1 := errors.New("1")
	err1 = xerror.WrapSkipf(1, err1, "2")
	err1 = xerror.WrapSkipf(1, err1, "3")
	assert.NotNil(t, err1, nil)
	assert.Equal(t, err1.Error(), "3: 2: 1")

	err2 := xerror.New("1")
	err2 = xerror.WrapSkipf(1, err2, "")
	assert.NotNil(t, err2, nil)
	assert.Equal(t, err2.Error(), "1")
}

func Test_Cause(t *testing.T) {
	err := errors.New("1")
	assert.Equal(t, xerror.Cause(err), err)

	err1 := errors.New("1")
	err1 = xerror.Wrap(err1, "2")
	err1 = xerror.Wrap(err1, "3")
	assert.Equal(t, xerror.Cause(err1).Error(), "1")

	err2 := xerror.New("1")
	assert.Equal(t, xerror.Cause(err2).Error(), "1")

	err3 := xerror.New("1")
	err3 = xerror.Wrap(err3, "2")
	err3 = xerror.Wrap(err3, "3")
	assert.Equal(t, xerror.Cause(err3).Error(), "1")
}

func Test_Format(t *testing.T) {
	err1 := errors.New("1")
	err1 = xerror.Wrap(err1, "2")
	err1 = xerror.Wrap(err1, "3")
	assert.NotNil(t, err1)
	assert.Equal(t, fmt.Sprintf("%s", err1), "3: 2: 1")
	assert.Equal(t, fmt.Sprintf("%v", err1), "3: 2: 1")

	err2 := xerror.New("1")
	err2 = xerror.Wrap(err2, "2")
	err2 = xerror.Wrap(err2, "3")
	assert.NotNil(t, err2, nil)
	assert.Equal(t, fmt.Sprintf("%-s", err2), "3")
	assert.Equal(t, fmt.Sprintf("%-v", err2), "3")
}

func Test_Stack(t *testing.T) {
	err := errors.New("1")
	assert.Equal(t, fmt.Sprintf("%+v", err), "1")

	err1 := errors.New("1")
	err1 = xerror.Wrap(err1, "2")
	err1 = xerror.Wrap(err1, "3")
	assert.NotNil(t, err1, nil)
	// fmt.Printf("%+v", err1)

	err2 := xerror.New("1")
	assert.NotNil(t, fmt.Sprintf("%+v", err2), "1")
	// fmt.Printf("%+v", err2)

	err3 := xerror.New("1")
	err3 = xerror.Wrap(err3, "2")
	err3 = xerror.Wrap(err3, "3")
	assert.NotNil(t, err3, nil)
	// fmt.Printf("%+v", err3)
}

func Test_Current(t *testing.T) {
	err := errors.New("1")
	err = xerror.Wrap(err, "2")
	err = xerror.Wrap(err, "3")
	assert.Equal(t, err.Error(), "3: 2: 1")
	assert.Equal(t, xerror.Current(err).Error(), "3")
}

func Test_Unwrap(t *testing.T) {
	err := errors.New("1")
	err = xerror.Wrap(err, "2")
	err = xerror.Wrap(err, "3")
	assert.Equal(t, err.Error(), "3: 2: 1")

	err = xerror.Unwrap(err)
	assert.Equal(t, err.Error(), "2: 1")

	err = xerror.Unwrap(err)
	assert.Equal(t, err.Error(), "1")

	err = xerror.Unwrap(err)
	assert.Nil(t, err)
}

func Test_Code(t *testing.T) {
	err1 := errors.New("123")
	assert.Equal(t, xerror.Code(err1).Code(), -1)
	assert.Equal(t, err1.Error(), "123")

	err2 := xerror.NewCode(xcode.CodeUnknown, "123")
	assert.Equal(t, xerror.Code(err2), xcode.CodeUnknown)
	assert.Equal(t, err2.Error(), "123")

	err3 := xerror.NewCodef(xcode.New(1, "", nil), "%s", "123")
	assert.Equal(t, xerror.Code(err3).Code(), 1)
	assert.Equal(t, err3.Error(), "123")

	err4 := xerror.NewCodeSkip(xcode.New(1, "", nil), 0, "123")
	assert.Equal(t, xerror.Code(err4).Code(), 1)
	assert.Equal(t, err4.Error(), "123")

	err5 := xerror.NewCodeSkipf(xcode.New(1, "", nil), 0, "%s", "123")
	assert.Equal(t, xerror.Code(err5).Code(), 1)
	assert.Equal(t, err5.Error(), "123")

	err6 := errors.New("1")
	err6 = xerror.Wrap(err6, "2")
	err6 = xerror.WrapCode(xcode.New(1, "", nil), err6, "3")
	assert.Equal(t, xerror.Code(err6).Code(), 1)
	assert.Equal(t, err6.Error(), "3: 2: 1")

	err7 := errors.New("1")
	err7 = xerror.Wrap(err7, "2")
	err7 = xerror.WrapCodef(xcode.New(1, "", nil), err7, "%s", "3")
	assert.Equal(t, xerror.Code(err7).Code(), 1)
	assert.Equal(t, err7.Error(), "3: 2: 1")

	err8 := errors.New("1")
	err8 = xerror.Wrap(err8, "2")
	err8 = xerror.WrapCodeSkip(xcode.New(1, "", nil), 100, err8, "3")
	assert.Equal(t, xerror.Code(err8).Code(), 1)
	assert.Equal(t, err8.Error(), "3: 2: 1")

	err9 := errors.New("1")
	err9 = xerror.Wrap(err9, "2")
	err9 = xerror.WrapCodeSkipf(xcode.New(1, "", nil), 100, err9, "%s", "3")
	assert.Equal(t, xerror.Code(err9).Code(), 1)
	assert.Equal(t, err9.Error(), "3: 2: 1")
}

func Test_SetCode(t *testing.T) {
	err := xerror.New("123")
	assert.Equal(t, xerror.Code(err).Code(), -1)
	assert.Equal(t, err.Error(), "123")

	err.(*xerror.Error).SetCode(xcode.CodeValidationFailed)
	assert.Equal(t, xerror.Code(err), xcode.CodeValidationFailed)
	assert.Equal(t, err.Error(), "123")
}

func Test_Json(t *testing.T) {
	err := xerror.Wrap(xerror.New("1"), "2")
	b, e := json.Marshal(err)
	assert.Equal(t, e, nil)
	assert.Equal(t, string(b), `"2: 1"`)
}

func Test_HasStack(t *testing.T) {
	err1 := errors.New("1")
	err2 := xerror.New("1")
	assert.Equal(t, xerror.HasStack(err1), false)
	assert.Equal(t, xerror.HasStack(err2), true)
}

func Test_Equal(t *testing.T) {
	err1 := errors.New("1")
	err2 := errors.New("1")
	err3 := xerror.New("1")
	err4 := xerror.New("4")
	assert.Equal(t, xerror.Equal(err1, err2), false)
	assert.Equal(t, xerror.Equal(err1, err3), true)
	assert.Equal(t, xerror.Equal(err2, err3), true)
	assert.Equal(t, xerror.Equal(err3, err4), false)
	assert.Equal(t, xerror.Equal(err1, err4), false)

}

func Test_Is(t *testing.T) {
	err1 := errors.New("1")
	err2 := xerror.Wrap(err1, "2")
	err2 = xerror.Wrap(err2, "3")
	assert.Equal(t, xerror.Is(err2, err1), true)
	err3 := xerror.Wrap(gorm.ErrRecordNotFound, "3")
	assert.Equal(t, xerror.Is(err3, gorm.ErrRecordNotFound), true)
	err4 := xerror.New(gorm.ErrRecordNotFound.Error())
	assert.Equal(t, xerror.Is(err4, gorm.ErrRecordNotFound), true)
	err5 := xerror.WrapCode(xcode.New(1, "record not found", ""), gorm.ErrRecordNotFound, "exes")
	assert.Equal(t, xerror.Is(err5, gorm.ErrRecordNotFound), true)
	t.Log((err5).(xerror.ICode).Code().Code())
}

func Test_HashError(t *testing.T) {
	err1 := errors.New("1")
	err2 := xerror.Wrap(err1, "2")
	err2 = xerror.Wrap(err2, "3")
	assert.Equal(t, xerror.HasError(err2, err1), true)
}

func Test_HashCode(t *testing.T) {
	err1 := errors.New("1")
	err2 := xerror.WrapCode(xcode.CodeNotAuthorized, err1, "2")
	err3 := xerror.Wrap(err2, "3")
	err4 := xerror.Wrap(err3, "4")
	assert.Equal(t, xerror.HasCode(err1, xcode.CodeNotAuthorized), false)
	assert.Equal(t, xerror.HasCode(err2, xcode.CodeNotAuthorized), true)
	assert.Equal(t, xerror.HasCode(err3, xcode.CodeNotAuthorized), true)
	assert.Equal(t, xerror.HasCode(err4, xcode.CodeNotAuthorized), true)
}
