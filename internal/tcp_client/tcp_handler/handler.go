package tcp_handler

import (
	"auto_healer/internal/auto"
	"auto_healer/internal/simulator"
	"context"
	"fmt"
	"net"
	"tcp_packet"
	"time"

	log "logger"

	"github.com/micmonay/keybd_event"
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
	log.Trace().Msgf("receiving tcp data [%d] from [%s]...", len(data), remoteAddr)

	pktIntf, err := tcp_packet.DeserializePacket(data)
	if err != nil {
		return fmt.Errorf("error at deserializing packet from [%s]: %s", remoteAddr, err.Error())
	}

	switch packet := pktIntf.(type) {
	case *tcp_packet.PacketPressed:
		{
			log.Trace().Msgf("received from [%s] packetType [0x%02X] inputData [0x%02X]", remoteAddr, packet.PacketType, packet.InputData)

			// 서버로부터 브로드캐스트 받은 키 입력에 따라 자동 동작 수행
			switch packet.InputData {
			case tcp_packet.KEY_F2: // F2: 력
				simulator.SendKeyboardInput(keybd_event.VK_0)

			case tcp_packet.KEY_F3: // 희원
				simulator.SendKeyboardInput(keybd_event.VK_5)

			case tcp_packet.KEY_F4: // F4: 희원첨
				simulator.SendKeyboardInput(keybd_event.VK_6)

			case tcp_packet.KEY_F5:
				// do nothing, will be use at the server to manual move...

			case tcp_packet.KEY_F6: // F6: Auto Move
				if auto.AutoMoveCtx == nil {
					auto.AutoMoveCtx, auto.AutoMoveCancel = context.WithCancelCause(context.Background())
					go auto.AutoMove(auto.AutoMoveCtx)
				} else {
					auto.AutoMoveCancel(fmt.Errorf("canceled by user"))
					auto.AutoMoveCtx = nil
				}

			case tcp_packet.KEY_F7: // F7: Auto Heal
				if auto.AutoHealCtx == nil {
					auto.AutoHealCtx, auto.AutoHealCancel = context.WithCancelCause(context.Background())
					go auto.AutoHeal(auto.AutoHealCtx)
				} else {
					auto.AutoHealCancel(fmt.Errorf("canceled by user"))
					auto.AutoHealCtx = nil
				}

			case tcp_packet.KEY_F8: // F8: Auto Debuf
				if auto.AutoDebufCtx == nil {
					auto.AutoDebufCtx, auto.AutoDebufCancel = context.WithCancelCause(context.Background())
					go auto.AutoDebuf(auto.AutoDebufCtx)
				} else {
					auto.AutoDebufCancel(fmt.Errorf("canceled by user"))
					auto.AutoDebufCtx = nil
				}

			case tcp_packet.KEY_F9: // F9: 시회 파혼 해독
				simulator.SendKeyboardInput(keybd_event.VK_7)
				simulator.SendKeyboardInput(keybd_event.VK_8)
				simulator.SendKeyboardInput(keybd_event.VK_9)

			// 방향키: Allow Manual Move
			case tcp_packet.KEY_UP:
				simulator.SendKeyboardInput(keybd_event.VK_UP)

			case tcp_packet.KEY_DOWN:
				simulator.SendKeyboardInput(keybd_event.VK_DOWN)

			case tcp_packet.KEY_LEFT:
				simulator.SendKeyboardInput(keybd_event.VK_LEFT)

			case tcp_packet.KEY_RIGHT:
				simulator.SendKeyboardInput(keybd_event.VK_RIGHT)
			}
		}

	case *tcp_packet.PacketBaramInfo:
		{
			hpPercent := packet.HpPercent
			mpPercent := packet.MpPercent

			log.Trace().Msgf("received from [%s] packetType [0x%02X] HP [%.1f%%] MP [%.1f%%]",
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
