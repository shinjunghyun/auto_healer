package tcp_handler

import (
	"fmt"
	"net"
	"tcp_packet"
	"time"

	log "logger"

	"github.com/dustin/go-humanize"
)

var (
	tcpConn net.Conn
)

func SetTcpConnection(conn net.Conn) {
	tcpConn = conn
}

func SendPacket(packet any) error {
	if tcpConn == nil {
		return fmt.Errorf("tcp connection is not set")
	}

	data, err := tcp_packet.SerializePacket(packet)
	if err != nil {
		return err
	}

	err = tcpConn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return fmt.Errorf("error setting write deadline: %s", err)
	}

	_, err = tcpConn.Write(data)

	return err
}

func Dispatcher(conn net.Conn, data []byte) error {
	remoteAddr := conn.RemoteAddr().String()
	log.Debug().Msgf("receiving tcp data [%d] from [%s]...", len(data), remoteAddr)

	pktIntf, err := tcp_packet.DeserializePacket(data)
	if err != nil {
		return fmt.Errorf("error at deserializing packet from [%s]: %s", remoteAddr, err.Error())
	}

	switch packet := pktIntf.(type) {
	case *tcp_packet.PacketPressed:
		{
			log.Info().Msgf("received from [%s] packetType [0x%02X] inputData [0x%02X]", remoteAddr, packet.PacketType, packet.InputData)
		}

	case *tcp_packet.PacketBaramInfo:
		{
			hpStr := humanize.Comma(int64(packet.HP))
			mpStr := humanize.Comma(int64(packet.MP))
			expStr := humanize.Comma(int64(packet.Exp))
			log.Info().Msgf("received from [%s] packetType [0x%02X] HP [%s] MP [%s] CoordX [%d] CoordY [%d] Exp [%s]",
				remoteAddr,
				packet.PacketType,
				hpStr,
				mpStr,
				packet.CoordX,
				packet.CoordY,
				expStr)
		}

	default:
		{
			log.Warn().Msgf("unhandled packet type from [%s]: %v", remoteAddr, packet)
		}
	}

	return nil
}
