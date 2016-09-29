package main

import (
	"net"

	"github.com/Sirupsen/logrus"
)

func main() {
	//Connect udp
	serverAddr, err := net.ResolveUDPAddr("udp", "192.168.7.10:514")
	localAddr, err := net.ResolveUDPAddr("udp", "192.168.100.10")
	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		logrus.Fatal("Error while parsing the logs", err)
	}
	defer conn.Close()

	//simple write
	conn.Write([]byte("Sep 18 10:28:52 orchestrator [programname][20310]: Log message "))
	conn.Write([]byte("Sep 18 10:28:52 192.168.100.10 [programname][20310]: Log message "))
}
