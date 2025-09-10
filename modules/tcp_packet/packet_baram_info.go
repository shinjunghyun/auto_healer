package tcp_packet

type PacketBaramInfo struct {
	PacketBase
	HP     uint32
	MP     uint32
	CoordX uint16
	CoordY uint16
	Exp    uint64
}
