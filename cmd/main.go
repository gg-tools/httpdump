package main

import (
	"github.com/gg-tools/httpdump/api/http"
	"github.com/gg-tools/httpdump/model/repository"
	"github.com/gg-tools/httpdump/model/service"
	"github.com/gg-tools/httpdump/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var (
	httpPort = utils.EnvInt("HTTP_PORT", 80)
	dataPath = utils.Env("DATA_PATH", "")
	logLevel = utils.Env("LOG_LEVEL", "info")
)

func main() {
	mainLogFile := openFile(filepath.Join(dataPath, "main.log"))
	defer mainLogFile.Close()
	accessLogFile := openFile(filepath.Join(dataPath, "access.log"))
	defer accessLogFile.Close()

	cacheRepo := repository.NewCache()
	dumper := service.NewDumper(cacheRepo)
	handler := http.NewHandler(dumper)

	configLog(mainLogFile, logLevel)
	http.ServerHTTP(httpPort, accessLogFile, handler.Route)
}

func openFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.WithError(err).Fatalf("failed to open file: %s", filePath)
	}
	return file
}

func configLog(file *os.File, level string) {
	if l, err := log.ParseLevel(level); err == nil {
		log.WithError(err).Errorf("parse level failed: level=%s", level)
		log.SetLevel(l)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339Nano})
}
