package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// errors
var (
	ErrSmallSize = errors.New("small size of data")
	ErrHeader    = "want: %s got: %s"
)

// Reader wav
type Reader struct {
	chunkID       []byte // 0...3 bytes
	chunkSize     uint32 // 4...7 bytes
	format        []byte // 8...11 bytes
	subchunk1ID   []byte // 12...15 bytes
	subchunk1Size uint32 // 16...19 bytes
	audioFormat   uint16 // 20...21 bytes
	numChannels   uint16 // 22...23 bytes
	sampleRate    uint32 // 24...27 bytes
	byteRate      uint32 // 28...31 bytes
	blockAlign    uint16 // 32...33 bytes
	bitsPerSample uint16 // 34...35 bytes
	subchunk2ID   []byte // 36...39 bytes
	subchunk2Size uint32 // 40...43 bytes
	reader        io.Reader
}

func (r *Reader) Read(data []byte) (n int, err error) {
	return r.reader.Read(data)
}

// GetNumChannels ...
func (r *Reader) GetNumChannels() uint16 {
	return r.numChannels
}

// GetSampleRate ...
func (r *Reader) GetSampleRate() uint32 {
	return r.sampleRate
}

// NewReader parse audio data in wav format
func NewReader(data []byte) (r *Reader, err error) {
	if len(data) < 44 {
		err = ErrSmallSize
		return
	}

	if bytes.Compare(data[0:4], tokenRiff) != 0 {
		err = fmt.Errorf(ErrHeader, tokenRiff, data[0:4])
		return
	}

	if bytes.Compare(data[8:12], tokenWAVE) != 0 {
		err = fmt.Errorf(ErrHeader, tokenRiff, data[0:4])
		return
	}

	if bytes.Compare(data[12:16], tokenFmt) != 0 {
		err = fmt.Errorf(ErrHeader, tokenRiff, data[0:4])
		return
	}

	r = &Reader{
		chunkID:       data[:4],
		chunkSize:     binary.LittleEndian.Uint32(data[4:8]),
		format:        data[8:12],  // 8...11
		subchunk1ID:   data[12:16], // 12...15
		subchunk1Size: binary.LittleEndian.Uint32(data[16:20]),
		audioFormat:   binary.LittleEndian.Uint16(data[20:22]),
		numChannels:   binary.LittleEndian.Uint16(data[22:24]),
		sampleRate:    binary.LittleEndian.Uint32(data[24:28]),
		byteRate:      binary.LittleEndian.Uint32(data[28:32]),
		blockAlign:    binary.LittleEndian.Uint16(data[32:34]),
		bitsPerSample: binary.LittleEndian.Uint16(data[34:36]),
		subchunk2ID:   data[36:40],
		subchunk2Size: binary.LittleEndian.Uint32(data[40:44]),
		reader:        bytes.NewReader(data[44:]),
	}
	return
}
