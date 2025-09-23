package auto

import (
	"context"
	"math/rand"
	"sync"
	"tcp_packet"
	"time"
)

const TabBoxCheckInterval = 1000 * time.Millisecond

var (
	mtx     sync.Mutex
	randObj = rand.New(rand.NewSource(time.Now().UnixNano()))

	AutoMoveCtx    context.Context
	AutoMoveCancel context.CancelCauseFunc

	AutoHealCtx    context.Context
	AutoHealCancel context.CancelCauseFunc

	AutoDebuffCtx    context.Context
	AutoDebuffCancel context.CancelCauseFunc

	ServerBaramInfoData ServerBaramInfo
	ClientBaramInfoData ClientBaramInfo

	lastTabBoxCheckAt time.Time

	ServerConfigInstance ConfigExternal
)

type ServerBaramInfo struct {
	tcp_packet.PacketBaramInfo
	LastUpdatedAt time.Time
}

type ClientBaramInfo struct {
	tcp_packet.PacketBaramInfo
	LastUpdatedAt time.Time
}
