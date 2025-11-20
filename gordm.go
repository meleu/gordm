package gordm

import (
	"math/rand"
	"time"
)

var Seed = time.Now().UnixNano()

// GenerateRandomInt returns a random integer between the min and max passed
// as arguments.
func GenerateRandomInt(min, max int) int {
	r := rand.New(rand.NewSource(Seed))
	return r.Intn(max-min+1) + min
}
