package main

import (
	"fmt"
	"log"
	"time"

	_ "net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	//"github.com/gavs/insert.go"
)

// device is "lo" means local
// we are created one go rest app, and runnig in local and we are doing
//packet here
//
// we  are getting payload in ip, tcp , application layer
var (
	//	device       string = "enp0s8"
	snapshot_len int32 = 65535
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = -1 * time.Second
	handle       *pcap.Handle
)

//var allPacketData []interface{}

type Packetdata struct {
	InterfaceName    interface{}
	PacketNo         interface{}
	Time             time.Time
	PacketLength     interface{}
	TcpPayloadLength interface{}
	// sourceMac net.HardwareAddr
	// destinationMac net.HardwareAddr
	// ethernetType layers.EthernetType
	//sourceMac interface{}
	// 	net.IP net.IP  layers.IPProtocol  uint8
	// uint32 uint32  uint8  uint16   []uint8
	//destinationMac   interface{}
	//ethernetType     interface{}
	SourceIP     interface{}
	DestnationIP interface{}
	Protocol     interface{}
	//version          interface{}
	SourcePort       interface{}
	DestinationPort  interface{}
	SequenceNumber   interface{}
	AcknowedgeNumber interface{}
	//dataOffset       interface{}
	//payload          interface{}
	//window interface{}
	// finFlag bool
	// synFlag bool
	// ackflag bool
}

// var headers packetdata

func PacketStream(c chan<- Packetdata, device string) {
	// Open device
	//headers = packetdata{"pc_no", "src_ip", "    des_ip", "    protocol", "  sp", "      dp", "         seq_no", "  akn_no"} //, "dataOffset", "window"}
	//allPacketData = append(allPacketData, headers)
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	// host 142.250.193.206
	// var filter string = "port 3306"

	// err = handle.SetBPFFilter(filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println("Only capturing TCP port 80 packets.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	i := 1
	for packet := range packetSource.Packets() {
		var packetData Packetdata
		// Do something with a packet here.
		fmt.Println("packet no:--", i)
		//fmt.Println(packet)

		// Iterate over all layers, printing out each layer type
		fmt.Println("All packet layers:")
		for _, layer := range packet.Layers() {
			fmt.Println("- ", layer.LayerType())
		}

		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			fmt.Println("%v**********************************************ETHER layer***********************************************", i)

			fmt.Println("Ethernet layer detected.")
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			Time := packet.Metadata().CaptureInfo.Timestamp
			PacketSize := packet.Metadata().CaptureInfo.CaptureLength
			fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
			fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
			// Ethernet type is typically IPv4 but could be ARP or other
			fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
			fmt.Println("Time:--", Time)
			fmt.Println("PacketSize:--", PacketSize)
			//fmt.Printf("%T %T  %T  %T ", i, ethernetPacket.SrcMAC, ethernetPacket.DstMAC, ethernetPacket.EthernetType)
			//el = map[string]interface{}{"packetNo": i, "sourceMac": ethernetPacket.SrcMAC , "destinationMac" : ethernetPacket.DstMAC , "ethernetType": ethernetPacket.EthernetType }

			fmt.Println("%v####################################################################################################", i)
			fmt.Println()
			packetData = Packetdata{InterfaceName: device, PacketNo: i, Time: Time, PacketLength: PacketSize}
			//packetData = Packetdata{packetNo: i}
		}

		// Let's see if the packet is IP (even though the ether type told us)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			fmt.Println("%v***********************IP layer**********************************************************************", i)
			fmt.Println("IPv4 layer detected.")
			ip, _ := ipLayer.(*layers.IPv4)
			// IP layer variables:
			// Version (Either 4 or 6)
			// IHL (IP Header Length in 32-bit words)
			// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
			// Checksum, SrcIP, DstIP
			fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
			fmt.Println("Protocol: ", ip.Protocol)
			fmt.Println("Version :- ", ip.Version)
			fmt.Println("IP Header Length in 32-bit words :- ", ip.IHL)
			fmt.Println("ip layer payload: ", ip.Payload)
			fmt.Println()
			//il = map[string]interface{}{"packetNo": i, "srcIP": ip.SrcIP, "dstIP": ip.DstIP,"protocol": ip.Protocol, "version": ip.Version}
			fmt.Printf("%T %T  %T  %T ", ip.SrcIP, ip.DstIP, ip.Protocol, ip.Version)

			// 	var  source_ip_address interface{}
			// 	source_ip_address = ip.SrcIP

			// 	var  destination_ip_address interface{}
			// 	destination_ip_address = ip.DstIP
			// destination_ip_address = fmt.Sprintf("%v", data.destnationIP)
			// 	protocol := fmt.Sprintf("%v", data.protocol)

			//packetData.sourceIP = ip.SrcIP
			var sip interface{}
			sip = ip.SrcIP
			packetData.SourceIP = fmt.Sprintf("%v", sip)
			//packetData.destnationIP = ip.DstIP
			var dip interface{}
			dip = ip.DstIP
			packetData.DestnationIP = fmt.Sprintf("%v", dip)
			var proto interface{}
			proto = ip.Protocol
			packetData.Protocol = fmt.Sprintf("%v", proto)
			//packetData.version = ip.Version
			fmt.Println("%v####################################################################################################", i)
		}

		// Let's see if the packet is IP (even though the ether type told us)
		// ipLayer := packet.Layer(layers.LayerTypeIPv6)
		// if ipLayer != nil {
		// 	fmt.Println("%v***********************IP layer*", i)
		// 	fmt.Println("IPv6 layer detected.")
		// 	ip, _ := ipLayer.(*layers.IPv6)
		// 	fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		// 	//fmt.Println("Protocol: ", ip.Protocol)
		// 	fmt.Println("ip layer payload: ", ip.Payload)

		// }

		tcplayer := packet.Layer(layers.LayerTypeTCP)
		if tcplayer != nil {
			fmt.Println("%v***********************TCP**********************************************************************", i)

			tcp, _ := tcplayer.(*layers.TCP)
			fmt.Println(tcp.SrcPort)
			fmt.Println(tcp.DstPort)
			// TCP layer variables:
			// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
			// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
			fmt.Println("Sequence number: ", tcp.Seq)
			fmt.Println("Ack: ", tcp.Ack)
			fmt.Println("DataOffset: ", tcp.DataOffset)
			fmt.Println("Window: ", tcp.Window)
			fmt.Println("FIN: ", tcp.FIN)
			fmt.Println("SYN: ", tcp.SYN)
			fmt.Println("ACK: ", tcp.ACK)
			fmt.Println(string(tcp.Payload))
			//tl = map[string]interface{}{"packetNo": i, "sequenceNo":  tcp.Seq, "ack":  tcp.Ack,"DataOffset": tcp.DataOffset, "SYN":  tcp.SYN, "ACK":tcp.ACK}
			fmt.Printf("%T %T  %T  %T  %T", tcp.Seq, tcp.Ack, tcp.DataOffset, tcp.Window, tcp.Payload)
			packetData.SequenceNumber = tcp.Seq
			packetData.AcknowedgeNumber = tcp.Ack
			packetData.TcpPayloadLength = len(tcp.Payload)
			//	packetData.dataOffset = tcp.DataOffset
			//	packetData.window = tcp.Window
			//	packetData.payload = tcp.Payload
			var sp, dp interface{}
			sp = tcp.SrcPort
			dp = tcp.DstPort
			packetData.SourcePort = fmt.Sprintf("%v", sp)
			packetData.DestinationPort = fmt.Sprintf("%v", dp)

			fmt.Println("%v#############################################################################################################################", i)
		}

		applicationLayer := packet.ApplicationLayer()
		if applicationLayer != nil {
			fmt.Println("%v***********************APPLIC*********************************************************************************************", i)

			fmt.Println("Application layer/Payload found.")
			fmt.Printf("Payload: %s\n ", applicationLayer.Payload())
			fmt.Println("%v######################################################################################################################################################", i)
		}

		fmt.Println("++++++++++++++++++++++++++++++++ALL PACKETS DATA ++++++++++++++++++++++++++++++++++++++")
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		//allPacketData = append(allPacketData, packetData)
		// for _, v := range allPacketData {
		// 		fmt.Println(v)
		// }
		c <- packetData
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		i++
		// Check for errors
		if err := packet.ErrorLayer(); err != nil {
			fmt.Println("Error decoding some part of the packet:", err)
		}
	}
}
