package config

type MysqlInfo struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	User   string `mapstructure:"user" json:"user"`
	Pwd    string `mapstructure:"pass_word" json:"pass_word"`
	DBname string `mapstructure:"db" json:"db"`
}

type DBConfig struct {
	Name      string    `mapstructure:"name" json:"name"`
	MysqlInfo MysqlInfo `mapstructure:"mysql" json:"mysql"`
}
