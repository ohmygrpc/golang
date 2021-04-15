package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Setting struct {
	ServiceName        string
	GRPCServerEndpoint string
	GRPCServerPort     string
	HTTPServerPort     string
	StatsdUDPPort      string

	Env       string
	SubEnvID  string
	Namespace string

	GracefulShutdownTimeoutMs int
}

func NewSetting() Setting {
	return Setting{
		ServiceName:        "golang",
		GRPCServerEndpoint: getEnv("GRPC_SERVER_ENDPOINT", "localhost:8080"),
		GRPCServerPort:     getEnv("GRPC_SERVER_PORT", "8080"),
		HTTPServerPort:     getEnv("HTTP_SERVER_PORT", "18080"),
		StatsdUDPPort:      getEnv("STATSD_UDP_PORT", "8125"),

		Env:       getEnv("ENV", "development"),
		SubEnvID:  getEnv("SUB_ENV_ID", "local"),
		Namespace: getEnv("NAMESPACE", "development-local"),

		GracefulShutdownTimeoutMs: mustAtoi(getEnv("GRACEFUL_SHUTDOWN_TIMEOUT_MS", "10000")),
	}
}

func MockSetting() Setting {
	return NewSetting()
}

func getEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		if defaultValue != "" {
			value = defaultValue
		} else {
			log.Fatalf("missing required environment variable: %v", key)
		}
	}
	return value
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}
