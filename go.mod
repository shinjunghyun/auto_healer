module auto_healer

go 1.24.7

replace (
	logger => ./modules/logger
	tcp_packet => ./modules/tcp_packet
)

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/kbinani/screenshot v0.0.0-20250624051815-089614a94018
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e
	github.com/micmonay/keybd_event v1.1.2
	github.com/moutend/go-hook v0.1.0
	golang.org/x/image v0.27.0
	golang.org/x/sys v0.36.0
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.9
	logger v0.0.0-00010101000000-000000000000
	tcp_packet v0.0.0-00010101000000-000000000000
)

require (
	github.com/gen2brain/shm v0.1.1 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)
