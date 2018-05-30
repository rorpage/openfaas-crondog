package main

import (
	"strconv"
	"time"
)

// HasEnv provides interface for os.Getenv
type HasEnv interface {
	Getenv(key string) string
}

// ReadConfig constitutes config from env variables
type ReadConfig struct {
}

func isBoolValueSet(val string) bool {
	return len(val) > 0
}

func parseBoolValue(val string) bool {
	if val == "true" {
		return true
	}
	return false
}

func parseIntOrDurationValue(val string, fallback time.Duration) time.Duration {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return time.Duration(parsedVal) * time.Second
		}
	}

	duration, durationErr := time.ParseDuration(val)
	if durationErr != nil {
		return fallback
	}
	return duration
}

func parseIntValue(val string, fallback int) int {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return parsedVal
		}
	}

	return fallback
}

// Read fetches config from environmental variables.
func (ReadConfig) Read(hasEnv HasEnv) CrondogConfig {
	cfg := CrondogConfig{
		writeDebug: false,
	}

	cfg.cronSchedule = hasEnv.Getenv("cron_schedule")
	cfg.functionURL = hasEnv.Getenv("function_url")
	cfg.functionData = hasEnv.Getenv("function_data")

	writeDebugEnv := hasEnv.Getenv("write_debug")
	if isBoolValueSet(writeDebugEnv) {
		cfg.writeDebug = parseBoolValue(writeDebugEnv)
	}

	return cfg
}

// CrondogConfig for the process
type CrondogConfig struct {

	// cron schedule to run
	cronSchedule string

	// function URL to invoke
	functionURL string

	// data to send to functionURL
	functionData string

	// writeDebug write console stdout statements to the container
	writeDebug bool
}
