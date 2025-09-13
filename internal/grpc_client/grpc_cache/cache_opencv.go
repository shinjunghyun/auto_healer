package grpc_cache

import (
	opencv_proto "auto_healer/external/proto/opencv-proto"
	"time"
)

const (
	FindTabBoxCacheDuration     = 300 * time.Millisecond
	GetHpMpPercentCacheDuration = 500 * time.Millisecond
)

var (
	FindTabBoxCacheData     FindTabBoxCache
	GetHpMpPercentCacheData GetHpMpPercentCache
)

type FindTabBoxCache struct {
	opencv_proto.FindTabBoxResponse
	CachedAt time.Time
}

type GetHpMpPercentCache struct {
	opencv_proto.GetHpMpPercentResponse
	CachedAt time.Time
}
