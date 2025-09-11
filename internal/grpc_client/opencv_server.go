package grpc_client

import (
	opencv_proto "auto_healer/external/proto/opencv-proto"
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

	if c.openCVServiceClient == nil {
		return nil, fmt.Errorf("opencv grpc client is not initialized, please call Connect() first")
	}

	res, err = c.openCVServiceClient.FindTabBox(ctx, req)
	return res, err
}
