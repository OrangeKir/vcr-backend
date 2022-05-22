package main

import (
	gateway_layer "Backend/gateway-layer"
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

	gateway_layer.AddPersonInfo(conn)
	gateway_layer.AddDivisionInfo(conn)
	gateway_layer.AddStaffInfo(conn)
	gateway_layer.AddPriceStructureInfo(conn)

	http.ListenAndServe("localhost:6080", nil)
}
