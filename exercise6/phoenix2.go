package main

import("net"
		"time"
		"encoding/json")

var Counter int
var Master int
var LastSignal time.Time

func connect() (*net.UDPConn){
	sendAddr, _ := net.ResolveUDPAddr("udp", "localhost:30011")
	conn, _ := net.DialUDP("udp", nil, sendAddr)
	return conn
}

func receive(conn *net.UDPConn){
    received := make([]byte,1)
	for ; true ; {
		_, _, _ = conn.ReadFromUDP(received)
		_ = json.Unmarshal(received, &Counter)
		println(Counter)
		LastSignal = time.Now()
	}
}

func send(conn *net.UDPConn, message int){
	a, _ := json.Marshal(message)
	println("Sender melding")
	_, _ = conn.Write(a)
}

func alive(){
	conn := connect()
	for ; true ; {
		send(conn, Counter)
		time.Sleep(100*time.Millisecond)
	}
}

func slave(){
	for ; true ; {
        //println(100*time.Millisecond)
        println(time.Since(LastSignal))
		if (time.Since(LastSignal)) > 200*time.Millisecond{
			println("Tar over som master!")
			Master = 1
			go alive()
			return
		}
	}
}

func main(){
	//Initialize
	Master = 0
	Counter = 0
	LastSignal = time.Now()

	// Create receiver-thread
	recAddr, _ := net.ResolveUDPAddr("udp",":30011")
    recConn, _ := net.ListenUDP("udp", recAddr)
	go receive(recConn)

	for ; true ; {
		if Master == 1{
			println(Counter)
			time.Sleep(1*time.Second)
			Counter = Counter + 1
			if Counter > 4{
				Counter = 0
			}
		} else {
			slave()
		}
	}
}
