package tcp_handler

import (
	"auto_healer/internal/auto"
	"fmt"
	"net"
	"tcp_packet"
	"time"

	log "logger"
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
			hpPercent := packet.HpPercent
			mpPercent := packet.MpPercent

			log.Info().Msgf("received from [%s] packetType [0x%02X] HP [%.1f%%] MP [%.1f%%]",
				remoteAddr,
				packet.PacketType,
				hpPercent,
				mpPercent)

			auto.ServerBaramInfoData = auto.ServerBaramInfo{
				PacketBaramInfo: *packet,
				LastUpdatedAt:   time.Now(),
			}
		}

	default:
		{
			log.Warn().Msgf("unhandled packet type from [%s]: %v", remoteAddr, packet)
		}
	}

	return nil
}
