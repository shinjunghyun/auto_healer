package auto

import (
	"sync"
	"tcp_packet"
	"time"
)

const TabBoxCheckInterval = 1000 * time.Millisecond

var (
	mtx sync.Mutex

	ServerBaramInfoData ServerBaramInfo
	ClientBaramInfoData ClientBaramInfo

	lastTabBoxCheckAt time.Time
)

type ServerBaramInfo struct {
	tcp_packet.PacketBaramInfo
	LastUpdatedAt time.Time
}

type ClientBaramInfo struct {
	tcp_packet.PacketBaramInfo
	LastUpdatedAt time.Time
}
