package main

import("net"
	   "strconv"
	   "strings")

type message struct{
	from string
	to string
	message string
	length int
}

func receive() (rec string){
	received := make([]byte,1024)
	recAddr, _ := net.ResolveUDPAddr("udp", ":30000")
	conn1, _ := net.ListenUDP("udp", recAddr)

	length, _, _ := conn1.ReadFromUDP(received)
	println(length)

	return string(received)
}

func send(s string){
	sendAddr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20015")
	conn, _ := net.DialUDP("udp", nil, sendAddr)

	b := []byte(s)
	
	rec, _ := conn.Write(b)
	println(rec)
}

func sendMessage(m *message){
	messageString := m.from + "+" + m.to + "+" + m.message + "+" + strconv.Itoa(m.length)
	println(messageString)

	//send(messageString)
}

func receiveMessage() (melding *message){
	mess := receive()
	//println(mess)

	//mess := "Data1+Data2+Heisann+1" 	
	
	m := new(message)

	text := strings.Split(mess, "+")
	m.from = text[0]
	m.to = text[1]
	m.message = text[2]
	m.length, _ = strconv.Atoi(text[3])

	return m
}

func printMessage(m *message){
	println("From: ", m.from)
	println("To: ", m.to)
	println("Message: ", m.message)
	println("Length: ", m.length, " word")
}

func main(){
	m := new(message)
	m.from = "Data1"
	m.to = "Data2"
	m.message = "Heisann"
	m.length = 1
}
