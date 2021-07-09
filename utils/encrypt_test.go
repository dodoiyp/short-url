package utils

import (
	"strconv"
	"testing"

	"github.com/lytics/base62"
)

func BenchmarkBase62Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base62.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(i), 10)))
	}
}
