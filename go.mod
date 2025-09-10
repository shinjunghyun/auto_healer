module auto_healer

go 1.24.7

replace (
	logger => ./modules/logger
	tcp_packet => ./modules/tcp_packet
)

require (
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/dustin/go-humanize v1.0.1
	github.com/kbinani/screenshot v0.0.0-20250624051815-089614a94018
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e
	github.com/micmonay/keybd_event v1.1.2
	github.com/moutend/go-hook v0.1.0
	golang.org/x/sys v0.36.0
	logger v0.0.0-00010101000000-000000000000
	tcp_packet v0.0.0-00010101000000-000000000000
)

require (
	github.com/gen2brain/shm v0.1.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
