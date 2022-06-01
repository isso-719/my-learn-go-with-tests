package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	excepted := "aaaaa"

	if repeated != excepted {
		t.Errorf("excepted %q but got %q", excepted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeat := Repeat("b")
	fmt.Println(repeat)
	// output: bbbbb
}
