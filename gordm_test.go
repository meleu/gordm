package gordm_test

import (
	"testing"

	"github.com/meleu/gordm"
)

func TestGenerateRandomInt(t *testing.T) {
	gordm.Seed = 1
	got := gordm.GenerateRandomInt(0, 32767)

	// I know it's 545 because I ran it with Seed = 1, and checked the result.
	want := 545

	if got != want {
		t.Errorf("\n got: %d\nwant: %d", got, want)
	}
}
