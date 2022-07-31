package config

// Server 服务器配置
type Server struct {
	Auth       Auth       `mapstructure:"auth" yaml:"auth"`             // 用户管理配置
	GRPC       GRPC       `mapstructure:"grpc" yaml:"grpc"`             // grpc配置
	HTTP       HTTP       `mapstructure:"http" yaml:"http"`             // http配置
	Databases  Databases  `mapstructure:"databases" yaml:"databases"`   // 数据库配置
	Dictionary Dictionary `mapstructure:"dictionary" yaml:"dictionary"` // 词典配置
}

// Auth 用户管理配置
type Auth struct {
	Username    string `mapstructure:"username" yaml:"username"`       // 用户名
	Password    string `mapstructure:"password" yaml:"password"`       // 密码
	EnableAdmin bool   `mapstructure:"enableAdmin" yaml:"enableAdmin"` // 是否开启认证
	Header      string `mapstructure:"header" yaml:"header"`           // 验证头
	Separator   string `mapstructure:"separator" yaml:"separator"`     // 分割符
}

// GRPC grpc配置
type GRPC struct {
	Addr string `mapstructure:"addr" yaml:"addr"` // grpc地址
}

// HTTP http配置
type HTTP struct {
	Addr string `mapstructure:"addr" yaml:"addr"` // HTTP地址
}

// Databases 数据库配置
type Databases struct {
	Shard   int32  `mapstructure:"shard" yaml:"shard"`     // 分片数量
	Path    string `mapstructure:"path" yaml:"path"`       // 数据目录
	Timeout int64  `mapstructure:"timeout" yaml:"timeout"` // 超时时间
	Debug   bool   `mapstructure:"debug" yaml:"debug"`     // 是否开启debug
}

// Dictionary 词典配置
type Dictionary struct {
	Path string `mapstructure:"path" yaml:"path"` // 词典目录
}
