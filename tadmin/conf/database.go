package conf

type Database struct {
	DBType       string `mapstructure:"dbType"`
	Host         string `mapstructure:"host"`         // 地址
	Port         int    `mapstructure:"port"`         // 端口
	Debug        bool   `mapstructure:"debug"`        // 是否打开Debug
	DbName       string `mapstructure:"DbName"`       // 数据库名
	User         string `mapstructure:"username"`     // 数据库用户名
	Password     string `mapstructure:"password"`     // 数据库密码
	MaxIdleConns int    `mapstructure:"MaxIdleConns"` //空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"MaxOpenConns"` //打开到数据库的最大连接数
	MaxLifetime  int    `mapstructure:"MaxLifetime"`  //连接可复用的最大时间
	TablePrefix  string `mapstructure:"TablePrefix"`
}
