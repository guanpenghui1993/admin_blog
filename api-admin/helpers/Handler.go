package helpers

import (
	"io"
	"os"
)

// gin access.log
func Access(logFile string) io.Writer {

	f, _ := os.Create(logFile)

	if ServerConfig.Mode == "debug" {
		return io.MultiWriter(f, os.Stdout)
	}

	return io.MultiWriter(f)
}
