package conf

type System struct {
	Env           string `mapstructure:"env"`            // 环境值
	Mode          string `mapstructure:"mode"`           // 运行模式
	Host          string `mapstructure:"host"`           // IP地址
	Port          int    `mapstructure:"port"`           // 端口号
	Stack         bool   `mapstructure:"stack"`          // 是否开启日志栈
	UseMultipoint bool   `mapstructure:"use-multipoint"` // 多点登录拦截
	Allowurl      string `mapstructure:"allowurl"`
	NoVerifyRoot  string `mapstructure:"noVerifyRoot"`
	NoVerify      string `mapstructure:"noVerify"`
	NoValidity    string `mapstructure:"noValidity"`
	Apisecret     string `mapstructure:"apisecret"`
}
