package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but expected %q", expected, repeated)
		}
	}

	t.Run("printing multiple 'a'", func(t *testing.T) {
		expected := Repeat("a", 5)
		repeated := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	//Outuput: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
