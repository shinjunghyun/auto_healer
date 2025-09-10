package cmd

import (
	"auto_healer/configs"
	"auto_healer/internal/auto/image_helper"
	"auto_healer/internal/auto/window_helper"
	"auto_healer/internal/config"
	"auto_healer/internal/helper"
	"auto_healer/internal/hooker"
	"auto_healer/internal/hooker/input_event_handler"
	"auto_healer/internal/tcp_client/tcp_handler"
	"image/png"
	log "logger"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

var (
	loggerConfig log.LoggerOptions
)

func init() {
	//logger config
	loggerConfig = config.DefaultLoggerConfigFromEnv()
	log.Init(loggerConfig)
}

func waitSignal(signals chan os.Signal) {
	s := <-signals
	log.Info().Msgf("got system signal: %v", s)
	shutdown()
}

func shutdown() {
	log.Info().Msgf("server is in shutting down...")
	os.Exit(0)
}

func AutoHealerStart(gitCommit, buildTime string) {
	log.Info().Msgf("start time: %s", time.Now().Local().Format(time.RFC3339))
	log.Info().Msgf("git commit: %s", gitCommit)
	log.Info().Msgf("build time: %s", buildTime)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go waitSignal(sigs)

	helper.ShowServicelogoPrint()

	go startTCPClient()
	go hooker.StartKeyboardHooker(input_event_handler.HandleInputEvent)

	// TODO: DELETE TEST CODE!!
	testCode()
	// TODO: END OF TEST CODE!!

	<-make(chan struct{})
}

func testCode() {
	user32 := windows.NewLazySystemDLL("user32.dll")
	isWindowVisible := user32.NewProc("IsWindowVisible")

	// Find the window with the title "PingInfoView"
	windowName := "MapleStory Worlds-바람의나라 클래식"
	hwnd, err := window_helper.FindWindow(windowName)
	if err != nil {
		log.Error().Msgf("Error finding window: %v\n", err)
		return
	}

	if hwnd == 0 {
		log.Error().Msgf("Window with title '%s' not found.\n", windowName)
		return
	}

	// Check if the window is visible
	visible, _, _ := isWindowVisible.Call(hwnd)
	if visible == 0 {
		log.Error().Msgf("Window with title '%s' is not visible.\n", windowName)
		return
	}

	// Capture the screen of the found window
	img, err := image_helper.CaptureScreen(hwnd)
	if err != nil {
		log.Error().Msgf("Failed to capture screen: %v\n", err)
		return
	}

	// Save the captured image to a file
	file, err := os.Create("./capture.png")
	if err != nil {
		log.Error().Msgf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Error().Msgf("Failed to save image: %v\n", err)
		return
	}

	log.Info().Msgf("Screen captured and saved to ./capture.png")
}

func startTCPClient() {
	// const host = "192.168.137.65"
	// const host = "127.0.0.1"
	const host = "49.172.185.152"
	const port = "9833"

	address := net.JoinHostPort(host, port)

	for {
		log.Info().Msgf("attempting to tcp connect to %s...", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Error().Msgf("failed to tcp connect: %v. retrying in %d seconds...", err, configs.TCP_RECONNECT_INTERVAL_SECONDS)
			time.Sleep(configs.TCP_RECONNECT_INTERVAL_SECONDS * time.Second)
			continue
		}

		log.Info().Msgf("connected to the tcp server [%s] successfully", conn.RemoteAddr().String())
		tcp_handler.SetTcpConnection(conn)

		handleConnection(conn)

		log.Warn().Msgf("disconnected to the tcp server [%s] retrying in %d seconds...", conn.RemoteAddr().String(), configs.TCP_RECONNECT_INTERVAL_SECONDS)
		time.Sleep(configs.TCP_RECONNECT_INTERVAL_SECONDS * time.Second)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 4096)

	for {
		n, err := conn.Read(buffer)
		if n == 0 {
			log.Info().Msgf("client [%s] has been disconnected", conn.RemoteAddr().String())
			return
		} else if err != nil {
			log.Error().Msgf("connection [%s] error: %s", conn.RemoteAddr().String(), err.Error())
			return
		}

		if err = tcp_handler.Dispatcher(conn, buffer[:n]); err != nil {
			log.Error().Msgf("error at dispatching data: %s", err.Error())
		}
	}
}
