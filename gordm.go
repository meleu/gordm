package gordm

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	Seed             = time.Now().UnixNano()
	RandomOrgBaseURL = "https://www.random.org"
)

// GenerateRandomInt returns a random integer between the given min and max.
func GenerateRandomInt(min, max int) int {
	r := rand.New(rand.NewSource(Seed))
	return r.Intn(max-min+1) + min
}

// GetRandomIntFromWeb returns a random integer, obtained from random.org,
// between the given min and max.
func GetRandomIntFromWeb(min, max int) int {
	// https://www.random.org/integers/?num=1&min=0&max=32767&col=1&base=10&format=plain
	url := fmt.Sprintf(
		"%s/integers/?num=1&min=%d&max=%d&col=1&base=10&format=plain",
		RandomOrgBaseURL, min, max,
	)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)
	respStr := strings.TrimSuffix(string(respBytes), "\n")
	number, _ := strconv.Atoi(respStr)

	return number
}
