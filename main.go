package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		getRandomScramble()
	}
	duration := time.Since(start)
	fmt.Println("Time taken for million scrambles: ", duration)
}
