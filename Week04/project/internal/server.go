package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"sunday960126/Go-001/Week04/project/internal/api"
	"sunday960126/Go-001/Week04/project/internal/config"
	"sunday960126/Go-001/Week04/project/internal/service"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	e := gin.Default()
	api.SetupRouter(e)
	return &Server{e}
}

func (s *Server) Start() {
	logrus.Infof("start to listen %s", config.Listen())
	if err := s.engine.Run(config.Listen()); err != nil {
		logrus.Panicf("listen error: %v", err)
	}
}

func InitServer() {
	db := initMysql()
	service.User = service.InitUserServer(db)

}

func initMysql() *gorm.DB {
	db, _ := gorm.Open("mysql", config.MySQLURL())
	//if err != nil {
	//	logrus.Panicf("mysql db connection error: %v", err)
	//}
	//logrus.Infof("connect to mysql successful!")
	//db.DB().SetMaxOpenConns(config.MySQLPoolSize())
	//
	//if logrus.GetLevel() >= logrus.DebugLevel {
	//	db.LogMode(true)
	//}
	return db
}
