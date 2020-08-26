package config

import "time"

// 关系型数据库配置信息
type Database struct {
	Type        string
	Host        string
	Port        uint32
	User        string
	Password    string
	Name        string
	TablePrefix string
}

// redis 配置信息
type Redis struct {
	Host        string
	Db          int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	MaxRetries  int
}

type Mongo struct {
	Host        string
	Database    string
	IdleTimeout time.Duration
}

type Server struct {
	RunMode      string
	HttpPort     uint32
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// 系统配置信息
type System struct {
	PageSize  uint32
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath string
	ImageMaxSize  uint32
	ImageAllowExt []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// 配置
type Config struct {
	Database Database
	Redis    Redis
	Mongo    Mongo
	System   System
	Server   Server
}

var Conf = &Config{}

func InitConfig(file string) {
	// 加载配置文件
	configLoader := NewLoader(file)
	configLoader.InitConfig()
}
