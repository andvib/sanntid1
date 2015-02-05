package main

import ("net")

func receive(){
	received := make([]byte,1024)
	recAddr, _ := net.ResolveUDPAddr("udp", ":30000")
	conn1, _ := net.ListenUDP("udp", recAddr)

	length, _, _ := conn1.ReadFromUDP(received)
	println(length)
	
	println(string(received))
}


func send(){
	sendAddr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20015")
	conn, _ := net.DialUDP("udp", nil, sendAddr)

	a := "KJHSIDHFKIJIURGIURSHIUGSHR&("
	b := []byte(a)
	
	rec, _ := conn.Write(b)
	println(rec)
}





func main(){
	send()
	receive()
}
