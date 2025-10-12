package main

import (
	"fmt"
	"math/rand"
)


func main() {
	ch := make(chan int)
	sqch := make(chan int)
	
	go func ()  {
		for i := 0; i < 10; i++{
			number := rand.Intn(100)
			ch <- number
		}
		close(ch)
	}()

	go func ()  {
		for n := range ch{
			sqch <- n * n
		}
		close(sqch)
	}()

	for sqn := range sqch{
		fmt.Println(sqn)
	}

}