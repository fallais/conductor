package routes

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
	"golang.org/x/net/ipv4"

	"github.com/fallais/conductor/shared"
	"github.com/fallais/conductor/utils"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// ScenarioController ...
type ScenarioController struct {
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewScenarioController ...
func NewScenarioController() *ScenarioController {
	return &ScenarioController{}
}

//------------------------------------------------------------------------------
// Protocol
//------------------------------------------------------------------------------

// ControllerError contains ...
type ControllerError struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//------------------------------------------------------------------------------
// Routes
//------------------------------------------------------------------------------

// List all Scenario
func (ctrl *ScenarioController) List(c web.C, w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, shared.Scenarii)
}

// Get a Message
func (ctrl *ScenarioController) Get(c web.C, w http.ResponseWriter, r *http.Request) {
	// Retrieve the ID
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		logrus.Errorln("Error:", err)
	}

	// Play all scenarii
	for key, scenario := range shared.Scenarii {
		if scenario.ID == id {
			logrus.Infoln("Playing the scenario :", key)
			for key, step := range scenario.Steps {
				logrus.Infoln("Playing the step :", key)
				err := playStep(step)
				if err != nil {
					logrus.Fatalln("Error:", err)
				}
			}
		}
	}

	// Pubish the response
	utils.JSONResponse(w, http.StatusOK, nil)
}

func playStep(step shared.Step) error {
	// Source
	srcIP := net.ParseIP(step.Events.LogSourceIP)

	// Destination
	dstIP := net.ParseIP("192.168.7.10")

	// IP
	ip := layers.IPv4{
		SrcIP:    srcIP.To4(),
		DstIP:    dstIP.To4(),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolUDP,
	}

	// UDP
	udp := layers.UDP{
		SrcPort: layers.UDPPort(50000),
		DstPort: layers.UDPPort(514),
	}

	// Options
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	// Checksum
	udp.SetNetworkLayerForChecksum(&ip)

	// Buffer
	buf := gopacket.NewSerializeBuffer()

	err := gopacket.SerializeLayers(buf, opts, &ip, &udp)
	if err != nil {
		return err
	}

	ipHeaderBuf := gopacket.NewSerializeBuffer()
	err = ip.SerializeTo(ipHeaderBuf, opts)
	if err != nil {
		return err
	}
	ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
	if err != nil {
		return err
	}

	// Send the log
	var packetConn net.PacketConn
	var rawConn *ipv4.RawConn
	packetConn, err = net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		return err
	}
	rawConn, err = ipv4.NewRawConn(packetConn)
	if err != nil {
		return err
	}

	for i := 0; i < step.Events.Nb; i++ {
		// Prepare the Payload
		payload := step.Events.Payload
		for key, value := range step.Events.Values {
			payload = strings.Replace(payload, fmt.Sprintf("{{%s}}", key), value[rand.Intn(len(value))], -1)
		}
		tcpPayloadBuf := gopacket.NewSerializeBuffer()
		err = gopacket.SerializeLayers(tcpPayloadBuf, opts, &udp, gopacket.Payload([]byte(payload)))
		if err != nil {
			return err
		}

		// Send the log
		err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil)
		if err != nil {
			return err
		}
		logrus.Infoln(fmt.Sprintf("packet of length %d sent!\n", (len(tcpPayloadBuf.Bytes()) + len(ipHeaderBuf.Bytes()))))
	}

	return nil
}
