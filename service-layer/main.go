package main

import (
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger := initZapLog()
	zap.ReplaceGlobals(logger)

	conn, err := AddDb()

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	http.ListenAndServe("localhost:6080", nil)
}
