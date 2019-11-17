package stat

import "github.com/shirou/gopsutil/net"

type Network struct {
	Counter struct {
		Name        string `json:"name"`        // interface name
		BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
		BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
		PacketsSent uint64 `json:"packetsSent"` // number of packets sent
		PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
		Errin       uint64 `json:"errin"`       // total number of errors while receiving
		Errout      uint64 `json:"errout"`      // total number of errors while sending
		Dropin      uint64 `json:"dropin"`      // total number of incoming packets which were dropped
		Dropout     uint64 `json:"dropout"`     // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
		Fifoin      uint64 `json:"fifoin"`      // total number of FIFO buffers errors while receiving
		Fifoout     uint64 `json:"fifoout"`     // total number of FIFO buffers errors while sending
	}
}

func NewNetwork() Network {
	n := net.Addr{}

	return Network{
		Counter: struct {
			Name        string `json:"name"`
			BytesSent   uint64 `json:"bytesSent"`
			BytesRecv   uint64 `json:"bytesRecv"`
			PacketsSent uint64 `json:"packetsSent"`
			PacketsRecv uint64 `json:"packetsRecv"`
			Errin       uint64 `json:"errin"`
			Errout      uint64 `json:"errout"`
			Dropin      uint64 `json:"dropin"`
			Dropout     uint64 `json:"dropout"`
			Fifoin      uint64 `json:"fifoin"`
			Fifoout     uint64 `json:"fifoout"`
		}{},
	}
}
