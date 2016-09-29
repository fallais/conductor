package main

import (
	"net"

	"github.com/Sirupsen/logrus"
)

func main() {
	// Prepare addresses
	qradarAddr, err := net.ResolveUDPAddr("udp", "192.168.7.10:514")
	proxyAddr, err := net.ResolveUDPAddr("udp", "192.168.6.30")

	// Open the connection
	conn, err := net.DialUDP("udp", proxyAddr, qradarAddr)
	if err != nil {
		logrus.Fatalln("Error while opening the connection", err)
	}
	defer conn.Close()

	// Send the logs
	for i := 0; i < 200; i++ {
		conn.Write([]byte("Sep 18 10:28:52 192.168.6.10 [programname][20310]: Log message "))
	}

}
