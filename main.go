package main

import (
	"fmt"
	"sync"
)

func main() {
	squareResultCh := make(chan int)
	doubleResultCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go calculateSquare(5, squareResultCh, &wg)

	wg.Add(1)
	go calculateDouble(7, doubleResultCh, &wg)

	go func() {
		wg.Wait()
		close(squareResultCh)
		close(doubleResultCh)
	}()

	for {
		select {
		case squareResult, ok := <-squareResultCh:
			if !ok {
				fmt.Println("Square goroutine finished.")
				squareResultCh = nil
			} else {
				fmt.Println("Square result:", squareResult)
			}
		case doubleResult, ok := <-doubleResultCh:
			if !ok {
				fmt.Println("Double goroutine finished.")
				doubleResultCh = nil
			} else {
				fmt.Println("Double result:", doubleResult)
			}
		}

		if squareResultCh == nil && doubleResultCh == nil {
			break
		}
	}
}

func calculateSquare(number int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := number * number
	resultCh <- result
}

func calculateDouble(number int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := number * 2
	resultCh <- result
}
