// Copyright union-pdd-go Author(https://houseme.github.io/union-pdd-go/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/union-pdd-go.

package util

import (
	"time"

	"go.uber.org/zap"
)

// Logger is a logger.
func Logger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", ""),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
