package scan

import (
	"getitle/src/pkg"
	"strings"
)

var sendData = pkg.Decode("YmXgZhZgYGDwYGBgYGRgYNghsAPEBbNB5JF/f2YmBEkL7D7NsIpB0aQKJMoS29H1Wuak4PwXHAzaAh4JTAwMDAAAAAD//w==")
var sendData2 = pkg.Decode("YmVgYBZgYGCQYGBgYGSAAVYGAAAAAP//")

// -default
func oxidScan(result *pkg.Result) {
	result.Port = "135"
	target := result.GetTarget()
	conn, err := pkg.TcpSocketConn(target, RunOpt.Delay)
	if err != nil {

		//fmt.Println(err)
		result.Error = err.Error()
		return
	}
	defer conn.Close()
	result.Open = true

	recv, err := pkg.SocketSend(conn, sendData, 4096)
	if err != nil {
		return
	}
	recv, err = pkg.SocketSend(conn, sendData2, 4096)
	if err != nil {
		return
	}
	recvStr := string(recv)
	if len(recvStr) < 42 {
		return
	}
	recvStr_v2 := recvStr[42:]
	packet_v2_end := strings.Index(recvStr_v2, "\x09\x00\xff\xff\x00\x00")
	if packet_v2_end == -1 {
		return
	}
	packet_v2 := recvStr_v2[:packet_v2_end]
	packet_v2 = strings.Replace(packet_v2, "\x00", "", -1)
	hostname_list := strings.Split(packet_v2, "\x07")

	result.Host = hostname_list[0]
	result.Title += pkg.AsciiEncode(strings.Join(hostname_list[1:], ","))
	result.HttpStat = "OXID"
	result.Port = "135 (oxid)"
	result.Protocol = "wmi"
	return
}
