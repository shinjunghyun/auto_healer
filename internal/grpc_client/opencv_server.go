package grpc_client

import (
	opencv_proto "auto_healer/external/proto/opencv-proto"
	"auto_healer/internal/grpc_client/grpc_cache"
	"context"
	"fmt"
	log "logger"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OpenCVClient struct {
	host                string
	port                uint16
	conn                *grpc.ClientConn
	openCVServiceClient opencv_proto.OpenCVServiceClient
}

func NewOpenCVServerClient(host string, port uint16) *OpenCVClient {
	return &OpenCVClient{
		host: host,
		port: port,
	}
}

func (c *OpenCVClient) Connect() error {
	log.Debug().Msgf("trying to connect to opencv-server grpc at [%s:%d]", c.host, c.port)
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", c.host, c.port), grpc.WithIdleTimeout(10*time.Second), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	c.conn = conn
	c.openCVServiceClient = opencv_proto.NewOpenCVServiceClient(conn)

	return nil
}

func (c *OpenCVClient) Close() error {
	if conn := c.conn; conn == nil {
		return nil
	} else if err := conn.Close(); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *OpenCVClient) FindTabBox(ctx context.Context, req *opencv_proto.FindTabBoxRequest) (res *opencv_proto.FindTabBoxResponse, err error) {
	log.Trace().Msgf("calling opencv-server grpc [FindTabBox]...")

	if time.Since(grpc_cache.FindTabBoxCacheData.CachedAt) < grpc_cache.FindTabBoxCacheDuration {
		log.Trace().Msgf("using cached FindTabBox data from %s ago", time.Since(grpc_cache.FindTabBoxCacheData.CachedAt).String())
		return &grpc_cache.FindTabBoxCacheData.FindTabBoxResponse, nil
	}

	if c.openCVServiceClient == nil {
		return nil, fmt.Errorf("opencv grpc client is not initialized, please call Connect() first")
	}

	res, err = c.openCVServiceClient.FindTabBox(ctx, req)
	return res, err
}

func (c *OpenCVClient) GetHpMpPercent(ctx context.Context, req *opencv_proto.GetHpMpPercentRequest) (res *opencv_proto.GetHpMpPercentResponse, err error) {
	log.Trace().Msgf("calling opencv-server grpc [GetHpMpPercent]...")

	if time.Since(grpc_cache.GetHpMpPercentCacheData.CachedAt) < grpc_cache.GetHpMpPercentCacheDuration {
		log.Trace().Msgf("using cached GetHpMpPercent data from %s ago", time.Since(grpc_cache.GetHpMpPercentCacheData.CachedAt).String())
		return &grpc_cache.GetHpMpPercentCacheData.GetHpMpPercentResponse, nil
	}

	if c.openCVServiceClient == nil {
		return nil, fmt.Errorf("opencv grpc client is not initialized, please call Connect() first")
	}

	res, err = c.openCVServiceClient.GetHpMpPercent(ctx, req)
	return res, err
}
