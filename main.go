package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}
	genElements := make([]int, size)
	src := rand.NewSource(time.Now().Unix())
	for i := 0; i < size; i++ {
		genElements[i] = rand.New(src).Int()
	}
	return genElements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 || data == nil {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	maxNum := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxNum {
			maxNum = data[i]
		}
	}
	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	maxNums := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(i, start, end int) {
			defer wg.Done()
			maxNums[i] = maximum(data[start:end])
		}(i, start, end) // добавил аргументы и параметры на случай, если версия go < 1.22
	}
	wg.Wait()
	return maximum(maxNums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	genElements := generateRandomElements(SIZE)
	if genElements == nil {
		fmt.Println("SIZE must be greater than zero")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	t := time.Now()
	maxNum := maximum(genElements)
	elapsed := time.Since(t).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	t = time.Now()
	maxNum = maxChunks(genElements)
	elapsed = time.Since(t).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNum, elapsed)
}
