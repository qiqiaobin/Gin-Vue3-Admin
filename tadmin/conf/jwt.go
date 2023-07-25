package conf

type JWT struct {
	SigningKey  string `mapstructure:"signing-key"`  // jwt签名
	ExpiresTime int64  `mapstructure:"expires_time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer_time"`  // 缓冲时间
	Issuer      string `mapstructure:"issuer"`       // 签发者
}
