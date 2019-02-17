package main

import (
	"net/http"
	"os"
	_ "net/http/pprof" // 东西没能用，也可以执行，代表用到其中一些东西
	"test/lang/errhandingnding/filelistingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 处理recover
		defer func() {
			if r := recover(); r != nil {
				log.Warn("Panic: %s", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		} ()

		err := handler(writer, request)

		if err != nil {
			log.Warn("Error handing request: %s", err.Error())

			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		panic(err)
	}
}