package main

import (
	gateway_layer "Backend/gateway-layer"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

const dbConnectionString = "postgres://postgres:1@localhost:5432/vcr"

func addDb() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbConnectionString)
	if err != nil {
		conn.Close(context.Background())
	}

	return conn, err
}

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()

	if err != nil {
		fmt.Printf(err.Error())
	}

	return logger
}

func main() {
	logger := initZapLog()
	zap.ReplaceGlobals(logger)

	conn, err := addDb()

	if err != nil {
		zap.L().Fatal(err.Error())
	}

	gateway_layer.AddPersonInfo(conn)
	gateway_layer.AddDivisionInfo(conn)
	gateway_layer.AddStaffInfo(conn)
	gateway_layer.AddPriceStructureInfo(conn)
	gateway_layer.AddPublicationsInfo(conn)
	gateway_layer.AddConferencesInfo(conn)

	http.ListenAndServe("localhost:6080", nil)
}
