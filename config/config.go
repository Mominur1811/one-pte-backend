package config

type GinMode string

const (
	DebugMode   = GinMode("debug")
	ReleaseMode = GinMode("release")
)

type DB struct {
	Host                string `json:"host"                    validate:"required"`
	Port                int    `json:"port"                    validate:"required"`
	Name                string `json:"name"                    validate:"required"`
	User                string `json:"user"                    validate:"required"`
	Pass                string `json:"pass"                    validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type Config struct {
	Mode             GinMode `json:"mode"                             validate:"required"`
	ServiceName      string  `json:"service_name"                     validate:"required"`
	HttpPort         int     `json:"http_port"                        validate:"required"`
	Db               DB      `json:"db"                               validate:"required"`
	MigrationSource  string  `json:"migrations"                       validate:"required"`
	RedisURL         string  `json:"redis_url"                        validate:"required"`
	HealthCheckRoute string  `json:"health_check_route"`
	JwtSecret        string  `json:"jwt_secret"`
}

var config *Config

func init() {
	config = &Config{}
}

func GetConfig() *Config {
	return config
}
