package configs

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataID    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type ServerConfig struct {
	Name           string         `mapstructure:"name" json:"name"`
	Host           string         `mapstructure:"host" json:"host"`
	Tags           []string       `mapstructure:"tags" json:"tags"`
	MysqlConfig    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	ConsulConfig   ConsulConfig   `mapstructure:"consul" json:"consul"`
	RocketMQConfig RocketMQConfig `mapstructure:"rocketmq" json:"rocketmq"`
	RedisConfig    RedisConfig    `mapstructure:"redis" json:"redis"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type RocketMQConfig struct {
	NameServer    []string `mapstructure:"name_server" json:"name_server"`
	ConsumerGroup string   `mapstructure:"consumer_group" json:"consumer_group"`
	Topic         string   `mapstructure:"topic" json:"topic"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
