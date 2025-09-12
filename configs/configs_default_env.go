package configs

const (
	SERVER_LIVE_TYPE string = "local"

	SERVICE_NAME string = "auto_healer"

	DEBUG_LEVEL      string = "DEBUG"
	LOG_STYLE        string = "CONSOLE"
	LOG_FILE_WRITE   bool   = false
	LOG_MAX_SIZE_MB  int    = 10
	LOG_MAX_BACKUPS  int    = 10
	LOG_MAX_AGE_DAYS int    = 14
	LOG_COMPRESSION  bool   = true

	TCP_RECONNECT_INTERVAL_SECONDS = 5

	TCP_SERVER_HOST = "49.172.185.152"
	TCP_SERVER_PORT = 9833

	// OPENCV_SERVER_HOST = "127.0.0.1"
	OPENCV_SERVER_HOST = "49.172.185.152"
	OPENCV_SERVER_PORT = 9834

	BARAM_WINDOW_TITLE  string = "MapleStory Worlds-바람의나라 클래식"
	BARAM_WINDOW_WIDTH  int    = 1382 // (바람영역 1024*768)
	BARAM_WINDOW_HEIGHT int    = 863  // (바람영역 1024*768)
)
