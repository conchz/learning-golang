package main

import (
	"fmt"
	"time"
)

func count(start int, end int, ch chan int) {
	c := 0
	for j := start; j < end; j++ {
		c = c + j
	}

	ch <- c
}

func main() {
	ts := time.Now().UnixNano()
	h := 1000000000
	sum := 0
	ch := make(chan int, 1000)
	numLength := cap(ch)

	for i := 0; i < numLength; i++ {
		num := h / numLength
		go count(num*i, num*i+num, ch)
	}

	for i := 0; i < numLength; i++ {
		select {
		case msg := <-ch:
			sum = sum + msg
		}
	}

	fmt.Printf("The sum of 1-1000000000 is %v, take %d milliseconds\n", sum+h, (time.Now().UnixNano()-ts)/1000000)
}
