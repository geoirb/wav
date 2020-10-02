package wav

import (
	"bytes"
	"errors"
	"io"
)

// errors
var (
	ErrSmallSize = errors.New("small size of data")
)

// NewReader audio wav
func NewReader(data []byte) (info *WAV, reader io.Reader, err error) {
	if len(data) < 44 {
		err = ErrSmallSize
		return
	}
	info = ParseWAV(data)
	reader = bytes.NewReader(data[44:])
	return
}
