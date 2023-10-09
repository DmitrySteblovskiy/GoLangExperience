package main

import "fmt"

func main() {
	ch := make(chan int)
	go say(ch)

	for a := range ch {
		fmt.Println(a)
	}
}

func say(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}

	// close(ch) - нужно добавить для решения дедлока
}
