package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C config

type config struct {
	Debug  bool   `yaml:"Debug"`
	Listen string `yaml:"Listen"`
	Env    string `yaml:"Env"`
	MySQL  mySQL  `yaml:"MySQL"`
}

func Init() {
	if err := viper.Unmarshal(&C); err != nil {
		logrus.Panicf("parse config file error: %s", err)
	}
	if C.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("config loaded: %+v", C)
}

func Listen() string {
	return C.Listen
}

type mySQL struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
	PoolSize int    `yaml:"PoolSize"`
}

func (m *mySQL) GetURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DBName)
}

func MySQLURL() string {
	return C.MySQL.GetURL()
}

func MySQLPoolSize() int {
	if C.MySQL.PoolSize <= 0 {
		return 100
	}
	return C.MySQL.PoolSize
}
