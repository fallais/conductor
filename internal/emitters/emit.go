package emitters

import (
	"net"

	"conductor/internal/models"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
)

// Emit ...
func Emit(messageCh <-chan *models.Message) error {
	packetConn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		return err
	}

	rawConn, err := ipv4.NewRawConn(packetConn)
	if err != nil {
		return err
	}

	for message := range messageCh {
		err := emit(rawConn, message)
		if err := nil {
			logrus.Errorln("error while sending the message")
		}
	}
}

func emit(rawConn *ipv4.RawConn, message *models.Message) error {
	// IPv4 layer (OSI 3)
	ipv4Layer := layers.IPv4{
		SrcIP:    message.SourceIP.To4(),
		DstIP:    message.DestinationIP.To4(),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolUDP,
	}

	// UDP layer (OSI 4)
	udpLayer := layers.UDP{
		SrcPort: layers.UDPPort(50000),
		DstPort: layers.UDPPort(514),
	}

	// Options
	opts := gopacket.SerializeOptions{
		//FixLengths:       true,
		//ComputeChecksums: true,
	}

	// UDP Checksum
	udpLayer.SetNetworkLayerForChecksum(&ipv4Layer)

	// Buffer
	buf := gopacket.NewSerializeBuffer()

	// Searialize
	err := gopacket.SerializeLayers(buf, opts, &ipv4Layer, &udpLayer, gopacket.Payload([]byte(message.Payload)))
	if err != nil {
		return err
	}

	ipHeaderBuf := gopacket.NewSerializeBuffer()
	err = ipv4.SerializeTo(ipHeaderBuf, opts)
	if err != nil {
		return err
	}
	ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
	if err != nil {
		return err
	}

	// Send the log
	err = rawConn.WriteTo(ipHeader, buf.Bytes(), nil)
	if err != nil {
		return err
	}

	return nil
}
