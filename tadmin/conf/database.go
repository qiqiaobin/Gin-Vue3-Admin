package conf

type Database struct {
	DBType       string `mapstructure:"dbType"`
	Host         string `mapstructure:"host"`           // 地址
	Port         int    `mapstructure:"port"`           // 端口
	Config       string `mapstructure:"config"`         // 高级配置
	DbName       string `mapstructure:"db-name"`        // 数据库名
	User         string `mapstructure:"username"`       // 数据库用户名
	Password     string `mapstructure:"password"`       // 数据库密码
	MaxIdleConns int    `mapstructure:"max_idle_conns"` //空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max_open_conns"` //打开到数据库的最大连接数
	MaxLifetime  int    `mapstructure:"max_life_time"`  //连接可复用的最大时间
	LogMode      string `mapstructure:"log-mode"`       // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap"`        // 是否通过zap写入日志文件
}
