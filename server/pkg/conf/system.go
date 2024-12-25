package conf

type System struct {
	Debug               bool
	Env                 string `mapstructure:"env" json:"env" yaml:"env"`                // 环境值
	Listen              string `mapstructure:"listen" json:"listen" yaml:"listen"`       // 监听地址
	DbType              string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType             string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix        string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	UseMultipoint       bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`                         // 多点登录拦截
	TokenExpiredTime    int64  `mapstructure:"token-expired-time" json:"token-expired-time" yaml:"token-expired-time"`             // 多点登录拦截
	TokenSecret         string `mapstructure:"token-secret" json:"token-secret" yaml:"token-secret"`                               // 多点登录拦截
	TokenOldExpiredTime int64  `mapstructure:"token-old-expired-time" json:"token-old-expired-time" yaml:"token-old-expired-time"` // 多点登录拦截
	ResourcePath        string `mapstructure:"resource-path" json:"resource-path" yaml:"resource-path"`                            // 资源文件路径
	ResourceDomain      string `mapstructure:"resource-domain" json:"resource-domain" yaml:"resource-domain"`                      // 资源文件域名
	BaseURI             string `mapstructure:"base-uri" json:"base-uri" yaml:"base-uri"`                                           // 基础URI
}
