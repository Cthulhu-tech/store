package random

import "math/rand"

func GetNumber(min int, max int) int {

	return rand.Intn(max-min) + min

}
