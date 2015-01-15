package main

import (."fmt"
	"runtime"
	"time")

var i int

func threadOne(){
	for x := 0 ; x <= 1000000 ; x++{
		i++
	}
}

func threadTwo(){
	for y := 0 ; y <= 1000000 ; y++{
		i--
	}
}

func main() {
	i = 0
	runtime.GOMAXPROCS(runtime.NumCPU())

	go threadOne()
	go threadTwo()

	time.Sleep(1000*time.Millisecond)
	Println(i)
}
