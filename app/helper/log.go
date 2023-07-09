package helper

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func EnsureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, err := os.Stat(dirName); err != nil {
		errCreate := os.MkdirAll(dirName, os.ModePerm)
		if errCreate != nil {
			panic(errCreate)
		}
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func WriteLog(filename string, param string) {
	path, _ := os.Getwd()
	currentTime := time.Now()

	// log to custom file
	LogFile := path + "/logs/" + currentTime.Format("2006-01-02") + "/" + filename

	EnsureDir(LogFile)

	// open log file
	logFile, err := os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println(param)
}

func RequestLog(filename string, ctx *fiber.Ctx) {
	type RequestData struct {
		Body   interface{}
		Param  interface{}
		Header interface{}
	}
	var body interface{}
	ctx.BodyParser(&body)
	requestData := RequestData{
		Body:   body,
		Param:  ctx.AllParams(),
		Header: ctx.GetReqHeaders(),
	}
	WriteLog(filename, "REQUEST ["+ctx.OriginalURL()+"] "+PrettyPrint(requestData))
}
