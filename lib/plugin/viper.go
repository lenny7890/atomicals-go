package plugin

import (
	"fmt"

	"github.com/spf13/viper"
)

// InitViperConfig init viper config
func InitViperConfig() *viper.Viper {
	var configFile string = "atomicals-config.toml"

	// flag.StringVar(&configFile, "c", configFileName, "config")
	// if len(configFile) == 0 {
	// 	panic("config file not exist！")
	// }
	// read config file
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file parse fail:%s\n", err))
	}

	return v

}
