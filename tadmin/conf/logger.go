// Package conf 日志配置
package conf

// LoggerConfig 日志配置
type Logger struct {
	Filename   string `toml:"filename"`    // 日志文件路径
	Level      string `toml:"level"`       // 日志级别: debug/info/warn/error/panic
	MaxSize    int    `toml:"max_size"`    // 日志文件旋转之前的最大大小
	MaxBackups int    `toml:"max_backups"` // 保留的旧日志文件的最大数量
	MaxAge     int    `toml:"max_age"`     // 保留旧日志文件的最大天数
	Color      bool   `toml:"color"`
	Caller     bool   `toml:"caller"`
	Compress   bool   `mapstructure:"compress"` // 是否压缩
}
