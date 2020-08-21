// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger
var err error

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("error initializing zap logger: %v", err)
	}
}

//LogInfo - logs information message to stdout
func LogInfo(name, message string) {
	defer logger.Sync()
	logger.Info(message,
		zap.Reflect("service-name:", name),
	)
}

//LogWarning - logs warning information message to stdout
func LogWarning(name, message string, err error) {
	defer logger.Sync()
	logger.Warn(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
}

//LogError - logs error message to stdout
func LogError(name, message string, err error) {
	defer logger.Sync()
	logger.Error(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
}

//LogPanic - logs error message to stdout and panics
func LogPanic(name, message string, err error) {
	defer logger.Sync()
	logger.Panic(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
}

//LogFatal - logs error message to stdout and panics and calls os.Exit(1)
func LogFatal(name, message string, err error) {
	defer logger.Sync()
	logger.Fatal(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
}
