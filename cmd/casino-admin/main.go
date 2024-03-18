package casinoadmin

import (
    "fmt"
    "math/rand"
)


func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	randomIndex := rand.Intn(len(numbers))
	fmt.Println(randomIndex)
}