package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

func Start() {
	fmt.Println("============================server start=====================================")
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("配置文件错误 %s", err.Error()))
	}
}
