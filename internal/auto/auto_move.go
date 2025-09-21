package auto

import (
	"auto_healer/internal/auto/baram_helper"
	"auto_healer/internal/simulator"
	"context"
	log "logger"
	"math/rand"
	"tcp_packet"
	"time"

	"github.com/micmonay/keybd_event"
)

var (
	lastClientCoordX, lastClientCoordY int32
	lastClientUpdateTime               time.Time

	randomMoveMilliseconds uint32 = 3000
)

func AutoMove(ctx context.Context) {
	ticker := time.NewTicker(80 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			log.Info().Msgf("auto move context is done")
			return

		case <-ticker.C:
			if false && isSelfHealing {
				log.Debug().Msgf("currently self-healing, will skip moving...")
				continue
			}
			if false && isDebuffing {
				log.Debug().Msgf("currently debuffing, will skip moving...")
				continue
			} else {
				// 클라이언트의 좌표 갱신
				if time.Since(ServerBaramInfoData.LastUpdatedAt) > 1*time.Second {
					log.Error().Msgf("received baram info is too old to use [%s] ago", time.Since(ServerBaramInfoData.LastUpdatedAt).String())
					continue
				} else if abs(int32(ClientBaramInfoData.X-ServerBaramInfoData.X)) <= 2 && abs(int32(ClientBaramInfoData.Y-ServerBaramInfoData.Y)) <= 2 {
					log.Trace().Msgf("client is already near the server position, will skip auto move: server (%d, %d) client (%d, %d)", ServerBaramInfoData.X, ServerBaramInfoData.Y, ClientBaramInfoData.X, ClientBaramInfoData.Y)
					continue
				} else {
					var err error

					// 클라이언트의 좌표 갱신
					ClientBaramInfoData.X, ClientBaramInfoData.Y, err = baram_helper.GetCoordinates()
					if err != nil {
						log.Error().Msgf("error at retrieving current coordinates, will skip auto move: %s", err.Error())
						continue
					}

					performAutoMove(ServerBaramInfoData.PacketBaramInfo, ClientBaramInfoData.PacketBaramInfo)
				}
			}
		}
	}
}

func performAutoMove(ServerCharacter, ClientCharacter tcp_packet.PacketBaramInfo) {
	var err error

	// log.Debug().Msgf("auto-move: server (%d, %d) client (%d, %d)", ServerCharacter.X, ServerCharacter.Y, ClientCharacter.X, ClientCharacter.Y)

	// 거리 계산
	xDistance := abs(int32(ClientCharacter.X) - int32(ServerCharacter.X))
	yDistance := abs(int32(ClientCharacter.Y) - int32(ServerCharacter.Y))

	// 벽 감지: 좌표 변화가 없으면 벽으로 판단 후 랜덤 이동
	if time.Since(lastClientUpdateTime) > time.Duration(randomMoveMilliseconds)*time.Millisecond &&
		ClientCharacter.X == int(lastClientCoordX) &&
		ClientCharacter.Y == int(lastClientCoordY) {
		log.Warn().Msg("wall detected! performing random move")
		randomMove()
		return
	}

	// 우선순위 결정 (기본적으로 더 먼 축을 우선)
	moveXFirst := xDistance > yDistance

	// 일정 확률로 우선순위 뒤바꾸기 (25% 확률)
	if randomChance(25) {
		moveXFirst = !moveXFirst
	}

	// 움직임 결정
	if moveXFirst {
		// X축 이동
		if ClientCharacter.X > ServerCharacter.X {
			moveLeft()
		} else if ClientCharacter.X < ServerCharacter.X {
			moveRight()
		}
	} else {
		// Y축 이동
		if ClientCharacter.Y > ServerCharacter.Y {
			moveUp()
		} else if ClientCharacter.Y < ServerCharacter.Y {
			moveDown()
		}
	}

	// 클라이언트의 좌표 갱신
	ClientCharacter.X, ClientCharacter.Y, err = baram_helper.GetCoordinates()
	if err != nil {
		log.Error().Msgf("error at retrieving current coordinates: %s", err.Error())
		return
	}

	// 좌표가 변경된 경우에만 last값 갱신
	if ClientCharacter.X != int(lastClientCoordX) || ClientCharacter.Y != int(lastClientCoordY) {
		lastClientCoordX = int32(ClientCharacter.X)
		lastClientCoordY = int32(ClientCharacter.Y)
		lastClientUpdateTime = time.Now()
	}
}

// 랜덤 이동 함수
func randomMove() {
	directions := []string{"left", "right", "up", "down"}
	randomDirection := directions[rand.Intn(len(directions))]

	switch randomDirection {
	case "left":
		moveLeft()
	case "right":
		moveRight()
	case "up":
		moveUp()
	case "down":
		moveDown()
	}
}

// 절대값 계산 함수
func abs(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

// 일정 확률로 true를 반환하는 함수
func randomChance(percent int) bool {
	return rand.Intn(100) < percent
}

func moveLeft() {
	mtx.Lock()
	defer mtx.Unlock()

	simulator.SendKeyboardInput(keybd_event.VK_LEFT)
}

func moveRight() {
	mtx.Lock()
	defer mtx.Unlock()

	simulator.SendKeyboardInput(keybd_event.VK_RIGHT)
}

func moveUp() {
	mtx.Lock()
	defer mtx.Unlock()

	simulator.SendKeyboardInput(keybd_event.VK_UP)
}

func moveDown() {
	mtx.Lock()
	defer mtx.Unlock()

	simulator.SendKeyboardInput(keybd_event.VK_DOWN)
}
