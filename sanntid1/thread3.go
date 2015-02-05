package main

import (."fmt"
	"runtime"
	"time")

var i int

var c = make(chan int,1)

func threadOne(){
	for x := 0 ; x <= 1000000 ; x++{
		for; <-c != 1; {}
		c<-0
		i++
		c<-1
	}
}

func threadTwo(){
	for y := 0 ; y <= 1000000 ; y++{
		for ;<-c != 1;{}
		c<-0
		i--
		c<-1
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
