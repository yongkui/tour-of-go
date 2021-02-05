package main

import "fmt"

func sum(s []int, c chan int) {
	fmt.Println("in-sum: c addr:", &c, s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("in-main: c addr:", &c)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x := <-c
	fmt.Println(x, &c)
	y := <-c
	fmt.Println(y, &c)
}

