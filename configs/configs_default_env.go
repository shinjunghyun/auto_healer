package configs

const (
	SERVER_LIVE_TYPE string = "local"

	SERVICE_NAME string = "auto_healer"

	DEBUG_LEVEL      string = "TRACE"
	LOG_STYLE        string = "CONSOLE"
	LOG_FILE_WRITE   bool   = false
	LOG_MAX_SIZE_MB  int    = 10
	LOG_MAX_BACKUPS  int    = 10
	LOG_MAX_AGE_DAYS int    = 14
	LOG_COMPRESSION  bool   = true

	TCP_RECONNECT_INTERVAL_SECONDS = 5

	TCP_SERVER_HOST = "49.172.185.152"
	TCP_SERVER_PORT = 9833

	BARAM_WINDOW_TITLE  string = "MapleStory Worlds-바람의나라 클래식"
	BARAM_WINDOW_WIDTH  int    = 800
	BARAM_WINDOW_HEIGHT int    = 600
)
