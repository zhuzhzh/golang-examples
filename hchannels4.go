package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	s0 = 0
)

func createcake(cs chan string, flavor string, count int) {
	for i := 0; i < count; i++ {
		cake := "making " + flavor + strconv.Itoa(i)
		cs <- cake
	}
}

func recievecake(cs chan string) {
	if s0 == 1 {
		for s := range cs {
			fmt.Println(s)
		}
	} else {
		for s := range cs {
			cake := <-cs
			fmt.Println(s)
			fmt.Println(cake)
		}
	}

}

func main() {
	cs := make(chan string)
	go createcake(cs, "strawberry", 5)
	go createcake(cs, "apple", 10)
	go recievecake(cs)

	du, _ := time.ParseDuration("2s")
	time.Sleep(du)

}
