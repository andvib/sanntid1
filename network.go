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

func connectMaster() (connection *net.UDPConn){
	sendAddr, err := net.ResolveUDPAddr("udp", "129.241.187.136:20015")
	conn, err2 := net.DialUDP("udp", nil, sendAddr)

	if (err != nil || err2 != nil){
		println("Kunne ikke koble til!")
		return nil
	}
	
	return conn
}

func receive(conn *net.UDPConn) (rec string){
	received := make([]byte,1024)
	recAddr, _ := net.ResolveUDPAddr("udp", ":30000")
	conn1, _ := net.ListenUDP("udp", recAddr)

	length, _, _ := conn1.ReadFromUDP(received)
	println(length)
	println(string(received))
	return string(received)
}

func send(s string, conn *net.UDPConn){
	/*sendAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20015")
	conn, err2 := net.DialUDP("udp", nil, sendAddr)

	if (err != nil){
		println("Kunne ikke koble til!")
		return
	}
	
	if (err2 != nil){
		println("Kunne ikke koble til!")
		return
	}*/

	b := []byte(s)
	
	rec, _ := conn.Write(b)
	println("hei!",rec)
}

func sendMessage(m *message, conn *net.UDPConn){
	messageString := m.from + "+" + m.to + "+" + m.message + "+" + strconv.Itoa(m.length)
	//println(messageString)

	send(messageString, conn)
}

func receiveMessage(conn *net.UDPConn) (melding *message){
	mess := receive(conn)
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

	/*x := true
	for ; x ; {
		sendMessage(m)
	}*/
	conn := connectMaster()
	sendMessage(m, conn)
	receive(conn)
}
