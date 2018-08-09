package random

import (
	"math/rand"
	"time"
)

//MakeNumber randomized number given min and max
func MakeNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
