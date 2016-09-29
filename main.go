package main

import (
	"flag"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/zenazn/goji"
)

func main() {
	//Connect udp
	conn, err := net.Dial("udp", "192.168.7.10:514")
	if err != nil {
		logrus.Fatal("Error while parsing the logs", err)
	}
	defer conn.Close()

	//simple Read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	//simple write
	conn.Write([]byte("Sep 18 10:28:52 server 2013-09-18T10:28:52Z server [programname][20310]: Log message "))

	flag.Set("bind", ":8080")
	goji.Serve()
}
