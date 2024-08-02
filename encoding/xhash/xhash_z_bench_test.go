// go test *.go -bench=".*"

package xhash_test

import (
	"github.com/mooncake9527/x/encoding/xhash"
	"testing"
)

var (
	str = []byte("This is the test string for hash.")
)

func Benchmark_BKDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.BKDR(str)
	}
}

func Benchmark_BKDR64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.BKDR64(str)
	}
}

func Benchmark_SDBM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.SDBM(str)
	}
}

func Benchmark_SDBM64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.SDBM64(str)
	}
}

func Benchmark_RS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.RS(str)
	}
}

func Benchmark_RS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.RS64(str)
	}
}

func Benchmark_JS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.JS(str)
	}
}

func Benchmark_JS64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.JS64(str)
	}
}

func Benchmark_PJW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.PJW(str)
	}
}

func Benchmark_PJW64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.PJW64(str)
	}
}

func Benchmark_ELF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.ELF(str)
	}
}

func Benchmark_ELF64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.ELF64(str)
	}
}

func Benchmark_DJB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.DJB(str)
	}
}

func Benchmark_DJB64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.DJB64(str)
	}
}

func Benchmark_AP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.AP(str)
	}
}

func Benchmark_AP64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xhash.AP64(str)
	}
}
