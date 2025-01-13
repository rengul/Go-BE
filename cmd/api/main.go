package main

import (
	"os"
	"re-home/config"
	"re-home/server"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.Println("Starting the application...")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
