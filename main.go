package main

import (
	"io"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/oguzhankuzlukluoglu/Microservices/config"
	"github.com/oguzhankuzlukluoglu/Microservices/models"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique;size:255"`
	Password string `gorm:"size:255"`
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

	
	db := models.GetDB()
	db.AutoMigrate(&User{})

	router := gin.Default()
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
