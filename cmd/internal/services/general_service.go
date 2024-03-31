package services

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle_weights(arr []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}
