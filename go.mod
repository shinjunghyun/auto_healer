module auto_healer

go 1.23.1

replace (
	logger => ./modules/logger
	tcp_packet => ./modules/tcp_packet
)

require (
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/micmonay/keybd_event v1.1.2
	github.com/moutend/go-hook v0.1.0
	logger v0.0.0-00010101000000-000000000000
	tcp_packet v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
