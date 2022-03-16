package utils

import (
	"fmt"
	"io"
)

// BsploggerOutput logger that can log to file (lumberjack) as well as stdout
type BsploggerOutput struct {
	writers []io.Writer
}

// NewLoggerOut combine all writers to create a new writer
func NewLoggerOut(writers ...io.Writer) BsploggerOutput {
	return BsploggerOutput{writers: writers}
}

func (logger *BsploggerOutput) Write(p []byte) (n int, err error) {
	var lastn int
	var lasterr error
	for _, writer := range logger.writers {
		lastn, lasterr = writer.Write(p)
	}

	if lasterr != nil {
		return lastn, fmt.Errorf("error from last log writer: %w", lasterr)
	}

	return lastn, nil
}
