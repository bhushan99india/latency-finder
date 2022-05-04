package main

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
)

func printTable(k string, v map[string]interface{}, table *simpletable.Table) {
	r := []*simpletable.Cell{
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", v["Interface_Name"])},
		{Text: k},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", v["total_tcp_length"].(int))},
		{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", v["total_time_taken"])},
	}
	table.Body.Cells = append(table.Body.Cells, r)
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}
func main() {

	result := make(map[string]map[string]interface{})
	var allPacketData []Packetdata
	//	var headers Packetdata
	// //, "dataOffset", "window"}
	//	headers = Packetdata{"interface_name", "pcket_no", "time", "length", 1, "src_ip", "    des_ip", "    protocol", "  sp", "      dp", "    seq_no", "  akn_no"}
	//	allPacketData = append(allPacketData, headers)

	c1 := make(chan Packetdata)

	// go PacketStream(c1, "enp0s8")
	// go PacketStream(c1, "enp0s3")
	// go PacketStream(c1, "lo")

	interfaces := []string{"enp0s8", "enp0s3", "lo"}

	for _, v := range interfaces {
		go PacketStream(c1, v)
	}
	for i := range c1 {
		allPacketData = append(allPacketData, i)
		for _, v := range allPacketData {
			var dummy interface{}
			if v.TcpPayloadLength == nil {
				dummy = 0
			} else {
				dummy = v.TcpPayloadLength
			}

			key := fmt.Sprintf("%v:%v->%v:%v", v.SourceIP, v.SourcePort, v.DestnationIP, v.DestinationPort)
			_, ok := result[key]
			if !ok {
				result[key] = map[string]interface{}{"Interface_Name": v.InterfaceName, "total_tcp_length": dummy.(int), "firsttime": v.Time, "time": v.Time}
			} else {
				old_tcp_length := result[key]["total_tcp_length"]
				firsttime := result[key]["firsttime"]
				diff := v.Time.Sub(firsttime.(time.Time))
				result[key]["total_tcp_length"] = old_tcp_length.(int) + dummy.(int)
				result[key]["time"] = v.Time
				result[key]["total_time_taken"] = diff
			}

			table := simpletable.New()

			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "INTERFACE NAME"},
					{Align: simpletable.AlignCenter, Text: "TRN ID"},
					{Align: simpletable.AlignCenter, Text: "DATA LENGTH"},
					{Align: simpletable.AlignCenter, Text: "LATENCY"},
				},
			}
			fmt.Println("start")
			for k, v := range result {
				printTable(k, v, table)
			}
			fmt.Println("end")
		}
	}
}
