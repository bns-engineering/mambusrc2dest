package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var dir string

func LogError(strlog string) {
	fmt.Println(strlog)
	log.Error(strlog)
}

func Log(strlog string) {
	fmt.Println(strlog)
	log.Info(strlog)
}

func LogFatal(strlog string) {
	fmt.Println(strlog)
	log.Fatal(strlog)
}

func LogWarning(strlog string) {
	fmt.Println(strlog)
	log.Warning(strlog)
}

func LogPanic(strlog string) {
	fmt.Println(strlog)
	log.Panic(strlog)
}

func OpenFileLog() {
	var filename string
	env := os.Getenv("ENV")
	filename = strings.Replace(filepath.Base(os.Args[0]), ".exe", "", 1)
	if len(env) == 0 {
		filename = filename + ".log"
	} else if env == "DEV" {
		filename = filename + "_" + GetStrDate() + ".log"
	} else if env == "STAGING" {
		filename = filename + "_" + GetStrDate() + ".log"
	} else if env == "PRODUCTION" {
		filename = filename + "_" + GetStrDate() + ".log"
	} else {
		log.Error("please define env")
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}

func GetStrDate() string {
	dt := time.Now().Format("2006-01-02")
	msg := fmt.Sprintf("%-24s", dt)
	return msg
}
