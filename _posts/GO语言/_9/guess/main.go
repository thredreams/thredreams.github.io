package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func guesser(guessans chan<- int, result <-chan int, wg *sync.WaitGroup) {
	top := 99
	bot := 0
	gues := (rand.Intn(100))
	guessans <- gues
	fmt.Printf("guesser: i guess %v!\n", gues)
	defer wg.Done()
	for v := range result {

		switch v {
		case 1:
			top = gues
			gues = (rand.Intn(top-bot) + bot)
			guessans <- gues
			fmt.Printf("guesser: i guess %v!\n", gues)
		case -1:
			bot = gues
			gues = (rand.Intn(top-bot) + bot)
			guessans <- gues
			fmt.Printf("guesser: i guess %v!\n", gues)
		case 0:
			// guessans <- 0
			fmt.Printf("guesser: i got it!\n")
			goto outer
		}
	}
outer:
}

func master(guessans <-chan int, result chan<- int, wg *sync.WaitGroup) {
	relans := rand.Intn(100)
	defer wg.Done()

	for v := range guessans {
		switch {
		case v > relans:
			result <- 1
			fmt.Printf("master: You guess bigger!\n")
		case v < relans:
			result <- -1
			fmt.Printf("master: You guess smaler!\n")
		case v == relans:
			result <- 0
			fmt.Printf("master: You got it!\n")
			goto outer
		}
	}
outer:
}
func main() {
	var wg sync.WaitGroup
	guessans := make(chan int)
	result := make(chan int)
	wg.Add(2)
	go master(guessans, result, &wg)
	go guesser(guessans, result, &wg)
	wg.Wait()
}
