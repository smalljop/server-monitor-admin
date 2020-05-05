package config

type Config struct {
	Mysql Mysql     `mapstructure:"mysql" json:"mysql"`
	Log   LogConfig `mapstructure:"log" json:"log"` //日志配置
	Jwt   Jwt       `mapstructure:"jwt" json:"jwt"`
}

//mysql
type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}

/**
日志配置
*/
type LogConfig struct {
	Prefix  string `mapstructure:"prefix" json:"prefix"`    //日志前缀
	LogFile bool   `mapstructure:"log-file" json:"logFile"` //是否生成文件
	Stdout  string `mapstructure:"stdout" json:"stdout"`    //输出级别
	File    string `mapstructure:"file" json:"file"`        //文件
}

//Jwt配置
type Jwt struct {
	ExpireTime int64  `mapstructure:"expire-time" json:"expireTime"` //过期时间
	Secret     string `mapstructure:"secret" json:"secret"`          //秘钥
	Issuer     string `mapstructure:"issuer" json:"issuer"`          //签发人
}
