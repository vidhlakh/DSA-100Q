package main

import (
	"testing"
)

func replaceAtIndex1(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

func replaceAtIndex2(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func generateString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

func BenchmarkSmall1(b *testing.B) {
	n := 10
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex1(str, replacement, index)
	}
}

func BenchmarkSmall2(b *testing.B) {
	n := 10
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex2(str, replacement, index)
	}
}

func BenchmarkMedium1(b *testing.B) {
	n := 100
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex1(str, replacement, index)
	}
}

func BenchmarkMedium2(b *testing.B) {
	n := 100
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex2(str, replacement, index)
	}
}

func BenchmarkBig1(b *testing.B) {
	n := 10000
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex1(str, replacement, index)
	}
}

func BenchmarkBig2(b *testing.B) {
	n := 10000
	str, index, replacement := generateString(n), n/2, 'B'

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceAtIndex2(str, replacement, index)
	}
}

func main() {
	var b *testing.B
	BenchmarkSmall1(b)
}
