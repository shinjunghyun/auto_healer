package auto

import (
	"sync"
	"tcp_packet"
	"time"
)

var (
	mtx sync.Mutex

	ServerBaramInfoData ServerBaramInfo
	ClientBaramInfoData tcp_packet.PacketBaramInfo
)

type ServerBaramInfo struct {
	tcp_packet.PacketBaramInfo
	LastUpdatedAt time.Time
}
