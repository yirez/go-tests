package whatarethepointofthese

import "testing"

func TestHello(t *testing.T) {
	want := "so, hi I guess"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
