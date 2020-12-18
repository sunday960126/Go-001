package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"sunday960126/Go-001/Week04/project/internal"
	"sunday960126/Go-001/Week04/project/internal/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func main() {
	cobra.OnInitialize(initServer)

	command := &cobra.Command{
		Use:   "app",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			internal.NewServer().Start()
		},
	}

	command.PersistentFlags().StringVarP(&cfgFile, "conf", "f", "etc/config.yaml", "config file (default is ./etc/config.yaml)")

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func initServer() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panicf("read config file error: %v", err)
	}
	config.Init()
	internal.InitServer()
}
