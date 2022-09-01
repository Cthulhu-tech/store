package random

import (
	"math/rand"
	"time"
)

func GetNumber(min int, max int) int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min

}
