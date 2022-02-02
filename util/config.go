package util

import (
	"fmt"

	"github.com/citymall/geo-loc/types"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type configType struct {
	data     types.Config
	isloaded bool
}

var config configType

func LoadConfig() {
	if !config.isloaded {
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
		viper.SetEnvPrefix("global")
		runmode := cast.ToString(viper.Get("runmode"))
		config.data.Mysql = viper.Get(runmode + ".mysql").(map[string]interface{})
		config.data.Mongodb = viper.Get(runmode + ".mongodb").(map[string]interface{})
		config.data.Redis = viper.Get(runmode + ".redis").(map[string]interface{})
		config.isloaded = true
	} else {
		//do nothig Just Chill !!!
	}
}

func GetConfig() types.Config {
	LoadConfig()
	return config.data
}
