package auto

import (
	"sync"
	"tcp_packet"
	"time"
)

const (
	WindowTitle = "MapleStory Worlds-바람의나라 클래식"
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
