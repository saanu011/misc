package main

import (
    "fmt"
    "strconv"
    "testing"
)

var smallInt = 40
var bigInt = 999999999999999

func Itoa(s int) string {
    i := strconv.Itoa(s)
    return i
}

func FormatInt(s int) string {
    i := strconv.FormatInt(int64(s), 10)
    return i
}

func Sprintf(s int) string {
    i := fmt.Sprintf("%d", s)
    return i
}

func BenchmarkItoa(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Itoa(smallInt)
    }
}

func BenchmarkFormatInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FormatInt(smallInt)
    }
}

func BenchmarkSprintf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sprintf(smallInt)
    }
}

func BenchmarkItoaBig(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Itoa(bigInt)
    }
}

func BenchmarkFormatIntBig(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FormatInt(bigInt)
    }
}

func BenchmarkSprintfBig(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sprintf(bigInt)
    }
}
