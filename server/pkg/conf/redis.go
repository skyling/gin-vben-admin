package conf

type Redis struct {
	Network  string `mapstructure:"network" json:"network" yaml:"network"`    // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	TaskDB   int    `mapstructure:"task-db" json:"task-db" yaml:"task-db"`    // redis的哪个数据库
	TokenDB  int    `mapstructure:"token-db" json:"token-db" yaml:"token-db"` // redis的哪个数据库
}
