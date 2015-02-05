package main

import(
	"net"
	. "fmt"
	"time"
	"runtime"
)

func connect()(conn *net.TCPConn){
	sendAddr, _ := net.ResolveTCPAddr("tcp", "129.241.187.136:33546")
	conn, _ =net.DialTCP("tcp", nil, sendAddr)
	
	a := "Conect to: 129.241.187.148:33546\x00"
	//b := []byte(a)
	send(a, conn)
	go tcp_listen(conn)
	
	//rec, _ := conn.Write(b)
	//println(rec)
		
	return

}

func send(a string, conn *net.TCPConn){
	println("inne i send")
	/*sendAddr, _ := net.ResolveTCPAddr("tcp", "129.241.187.136:33546")
	conn, _ :=net.DialTCP("tcp", nil, sendAddr)
	
	a := "Connect to: 129.241.187.148:33546\x00"*/
	b := []byte(a)

	//go tcp_listen(conn)
	
	rec, _ := conn.Write(b)
	println("rec: ",rec)
	/*for {
		rec, _ := conn.Write(b)
		println(rec)
		time.Sleep(1*time.Second)	
	}*/
}

func tcp_listen(sock *net.TCPConn){
	buf := make([]byte, 1024)
	for {
		sock.Read(buf)
		Println(string(buf))
		//time.Sleep(1*time.Second)
	}
}

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())
	conn := connect()
	println("Ute av connect")
	time.Sleep(1*time.Second)
	send("ouh", conn)
	time.Sleep(1*time.Second)
	// receive()
}
