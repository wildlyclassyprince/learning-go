package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertionCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("single character repeated N times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertionCorrectMessage(t, expected, repeated)
	})

	t.Run("word repeated N times", func(t *testing.T) {
		repeated := Repeat("smile", 2)
		expected := "smilesmile"
		assertionCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	syllables := []string{"ba", Repeat("na", 2)}
	fmt.Println(strings.Join(syllables, ""))
	//Output: banana
}
