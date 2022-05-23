package main

import (
	gateway_layer "Backend/gateway-layer"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const dbConnectionString = "postgres://dduzpxtujiiien:2bab6d135fa16753e26be65c4596c000e9e211c011fa92f72509bc0ddd4a46ba@ec2-176-34-211-0.eu-west-1.compute.amazonaws.com:5432/dcreheg4vihq8r"

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
}
