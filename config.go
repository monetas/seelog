// Copyright 2011 Cloud Instruments Co. Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sealog

import (
	"bytes"
	. "github.com/cihub/sealog/common"
	"github.com/cihub/sealog/config"
	"github.com/cihub/sealog/dispatchers"
	"github.com/cihub/sealog/format"
	"io"
	"os"
)

// LoggerFromConfigAsFile creates logger with config from file. File should contain valid sealog xml.
func LoggerFromConfigAsFile(fileName string) (LoggerInterface, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	conf, err := config.ConfigFromReader(file)
	if err != nil {
		return nil, err
	}
	
	return createLoggerFromConfig(conf)
}

// LoggerFromConfigAsBytes creates a logger with config from bytes stream. Bytes should contain valid sealog xml.
func LoggerFromConfigAsBytes(data []byte) (LoggerInterface, error) {
	conf, err := config.ConfigFromReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	
	return createLoggerFromConfig(conf)
}

// LoggerFromConfigAsString creates a logger with config from a string. String should contain valid sealog xml.
func LoggerFromConfigAsString(data string) (LoggerInterface, error) {
	return LoggerFromConfigAsBytes([]byte(data))
}

// LoggerFromWriterWithMinLevel creates a simple logger for usage with non-Sealog systems. 
// Creates logger that writes to output with minimal level = minLevel.
func LoggerFromWriterWithMinLevel(output io.Writer, minLevel LogLevel) (LoggerInterface, error) {
	constraints, err := NewMinMaxConstraints(minLevel, CriticalLvl)
	if err != nil {
		return nil, err
	}

	dispatcher, err := dispatchers.NewSplitDispatcher(format.DefaultFormatter, []interface{}{output})
	if err != nil {
		return nil, err
	}

	conf, err := config.NewConfig(constraints, make([]*LogLevelException, 0), dispatcher, config.SyncLoggerType, nil)
	if err != nil {
		return nil, err
	}
	
	return createLoggerFromConfig(conf)
}
