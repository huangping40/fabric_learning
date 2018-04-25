/**
https://gist.github.com/dtjm/c6ebc86abe7515c988ec
go test -bench=. -benchmem
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	testData = []string{"a", "b", "c", "d", "e"}
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		_ = s
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" + testData[1] + ":" + testData[2] + ":" + testData[3] + ":" + testData[4]
		_ = s
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.WriteString(testData[0])
		b.WriteByte(':')
		b.WriteString(testData[1])
		b.WriteByte(':')
		b.WriteString(testData[2])
		b.WriteByte(':')
		b.WriteString(testData[3])
		b.WriteByte(':')
		b.WriteString(testData[4])
		s := b.String()
		_ = s
	}
}
