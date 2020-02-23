package main

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expexted := "aaaaa"

	if repeated != expexted {
		t.Errorf("expected '%q' but not '%q'", expexted, repeated)
	}
}

func ExampleRepeat() {
	out := Repeat("t")
	fmt.Println(out)
	// Output:
	// ttttt
}
