package initialization

import (
	"github.com/spf13/viper"

	"github.com/ahlemarg/BMI/src/globals"
)

func GetEnv(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func init() {
	res := GetEnv("DEBUG_MODLE")

	config_file := "src/db.yaml"
	if res {
		config_file = "src/db.yaml"
	}

	v := viper.New()

	v.SetConfigFile(config_file)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&globals.DBconfig); err != nil {
		panic(err)
	}
}
