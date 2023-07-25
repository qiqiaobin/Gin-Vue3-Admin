package conf

type Server struct {
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Logger   Logger   `mapstructure:"log"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Database Database `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha  Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors     CORS     `mapstructure:"cors" json:"cors" yaml:"cors"` // 跨域配置
}
