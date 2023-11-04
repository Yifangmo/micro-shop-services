package configs

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type ServerConfig struct {
	Name             string         `mapstructure:"name" json:"name"`
	Host             string         `mapstructure:"host" json:"host"`
	Tags             []string       `mapstructure:"tags" json:"tags"`
	ConsulConfig     ConsulConfig   `mapstructure:"consul" json:"consul"`
	MysqlConfig      MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	GoodsSrvName     string         `mapstructure:"goods_srv_name" json:"goods_srv_name"`
	InventorySrvName string         `mapstructure:"inventory_srv_name" json:"inventory_srv_name"`
	RocketMQConfig   RocketMQConfig `mapstructure:"rocketmq" json:"rocketmq"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	DB       string `mapstructure:"db" json:"db"`
}

type RocketMQConfig struct {
	NameServer             []string `mapstructure:"name_server" json:"name_server"`
	ConsumerGroup          string   `mapstructure:"consumer_group" json:"consumer_group"`
	GiveBackInventoryTopic string   `mapstructure:"topic" json:"topic"`
}
