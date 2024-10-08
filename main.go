package main

import (
	"io"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/oguzhankuzlukluoglu/Microservices/config"
	"github.com/oguzhankuzlukluoglu/Microservices/models"
	"github.com/oguzhankuzlukluoglu/Microservices/routes"
	"github.com/sirupsen/logrus"
)
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	runmod := "prod"
	configPath := "/config/config.json"
	if user.Name != "root" {
		configPath = "./config/config.json"
		runmod = "dev"
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	initLogger()
	config.SetConfig("runmod", runmod)
	config.SetConfig("configs", configPath)

	config.LoadConfig(runmod)

	// Veritabanı bağlantısını başlatır
	models.SetDB(config.GetConnectionString())

	router := gin.Default()
	routes.RegisterTaskRoutes(router)
	if err := router.Run(":8000"); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(colorable.NewColorableStdout())
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
