package gordm_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/meleu/gordm"
)

func TestGenerateRandomInt(t *testing.T) {
	gordm.Seed = 1
	got := gordm.GenerateRandomInt(0, 32767)

	// I know it's 545 because I ran it with Seed = 1, and checked the result.
	want := 545
	assertEqualInt(t, got, want)
}

func TestGetRandomIntFromWeb(t *testing.T) {
	min, max, want := 0, 100, 42

	fakeServer := newFakeRandomServer(want)
	gordm.RandomOrgBaseURL = fakeServer.URL
	got := gordm.GetRandomIntFromWeb(min, max)

	assertEqualInt(t, got, want)
}

// ----------------------------------------------------------------------
// Test helper functions
// ----------------------------------------------------------------------

// newFakeRandomServer creates a mocked server that responds with the
// integer given as argument.
//
// useful doc: https://pkg.go.dev/net/http/httptest
func newFakeRandomServer(num int) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			numStr := strconv.Itoa(num)
			w.Write([]byte(numStr))
		}),
	)
}

func assertEqualInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("\n got: %d\nwant: %d", got, want)
	}
}
