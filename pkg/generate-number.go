package pkg

import "math/rand"

func GenerateRandomNumber(length int) int {

	min := int64(1)
	max := int64(1)
	for i := 1; i < length; i++ {
		min *= 10
		max *= 10
	}
	max = max*10 - 1

	return rand.Intn(int(max-min+1)) + int(min)
}
