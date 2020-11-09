package config

const (
	AppMode = "local"
	AppPort = "127.0.0.1:8000"
	AppName = "gin-api"

	// 签名超时时间 s
	SignExpiry = 120

	AppDebug = false
	// 以 根目录main.go为根
	LogPath = "/storage/logs/"
)
