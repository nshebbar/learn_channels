package main

import "fmt"

func main() {
	
	in := make(chan string)
	out := make(chan string)

	in3 := make(chan int, 10)
	out3 := make(chan int)

	out1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out1 <- localI * 2
		}(i)
	}
	var result []int
	for i := 0; i < 10; i++ {
		val := <-out1
		result = append(result, val)
	}
	fmt.Println(result)

	go func() {
		name := <-in
		out <- fmt.Sprintf("Hello, " + name)
	}()
	in <- "Bob"
	close(in)
	message := <-out
	fmt.Println(message)

	go func() {
		for {
			i := <- in3
			out3 <- i * 2
		}
	}()

	in3 <- 1
	o1 := <- out3
	in3 <- 2
	o2 := <- out3
	fmt.Println(o1, o2)

	for i := 0; i < 10; i++ {
		in3 <- i
	}
	close(in3)

	go func() {
		for {
			i, ok := <-in3
			if !ok {
				close(out3)
				break
			}
			out3 <- i * 2
		}
	}()
	
	for v := range out3 {
		fmt.Println(v)
	}
}
