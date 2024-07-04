package main

import (
	"fmt"
	"math/rand"
	"time"
)

func weightedRandomChoice(values []int, weights []float64) int {
    if len(values) != len(weights) {
        panic("Values and weights must be the same length")
    }

    // Инвертируем веса
    invWeights := make([]float64, len(weights))
    for i, w := range weights {
        invWeights[i] = 1.0 / w
    }

    // Составляем список кумулятивных сумм
    cumSum := make([]float64, len(invWeights))
    cumSum[0] = invWeights[0]
    for i := 1; i < len(invWeights); i++ {
        cumSum[i] = cumSum[i-1] + invWeights[i]
    }

    // Генерируем случайное число
    rand.Seed(time.Now().UnixNano())
    r := rand.Float64() * cumSum[len(cumSum)-1]

    // Находим элемент, соответствующий случайному числу
    for i, cs := range cumSum {
        if r < cs {
            return values[i]
        }
    }

    return values[len(values)-1] // На случай, если не нашли (что маловероятно)
}

func main() {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    weights := []float64{10, 50, 100, 150, 200, 250, 300, 350, 450, 500} 

    result := weightedRandomChoice(values, weights)
    fmt.Println("Выбранный элемент:", result)
}