package util

import (
	"github.com/op/go-logging"
	"os"
	"runtime"
	"strings"
)

func GetLogger() *logging.Logger {

	pc, filename, line, ok := runtime.Caller(1)
	if !ok {
		panic("runtime.Caller() failed")
	}
	filepath := "tmp/" + split(filename) + ".log"

	var log = logging.MustGetLogger("logger")
	// prepare logging
	var format = logging.MustStringFormatter(`[%{level:-8s}] %{time:2006-01-02 15:04:05} | %{callpath:2} | %{message}`)

	file, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFile := logging.NewLogBackend(file, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFile, backendFormatter)

	log.Infof("Logger %v initialized", split(filename))
	log.Infof("file: %v line #%d, func: %v \n", split(filename)+".go", line, runtime.FuncForPC(pc).Name())
	return log
}

func split(any string) string {
	res := strings.Split(any, "/")
	res = strings.Split(res[len(res)-1], ".")
	return res[0]
}
