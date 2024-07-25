package xerror_test

import (
	"errors"
	"github.com/mooncake9527/x/xerrors/xcode"
	"github.com/mooncake9527/x/xerrors/xerror"
	"testing"
)

var (
	// Wrap* 函数基准测试的基本错误。
	baseError = errors.New("test")
)

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.New("test")
	}
}

func Benchmark_Newf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.Newf("%s", "test")
	}
}

func Benchmark_Wrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.Wrap(baseError, "test")
	}
}

func Benchmark_Wrapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.Wrapf(baseError, "%s", "test")
	}
}

func Benchmark_NewSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewSkip(1, "test")
	}
}

func Benchmark_NewSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewSkipf(1, "%s", "test")
	}
}

func Benchmark_NewCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewCode(xcode.New(500, "", nil), "test")
	}
}

func Benchmark_NewCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewCodef(xcode.New(500, "", nil), "%s", "test")
	}
}

func Benchmark_NewCodeSkip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewCodeSkip(xcode.New(1, "", nil), 500, "test")
	}
}

func Benchmark_NewCodeSkipf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.NewCodeSkipf(xcode.New(1, "", nil), 500, "%s", "test")
	}
}

func Benchmark_WrapCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.WrapCode(xcode.New(500, "", nil), baseError, "test")
	}
}

func Benchmark_WrapCodef(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xerror.WrapCodef(xcode.New(500, "", nil), baseError, "test")
	}
}
