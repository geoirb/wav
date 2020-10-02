package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// errors
var (
	ErrSmallSize = errors.New("small size of data")
	ErrFormat    = errors.New("wav format")
)

// Reader wav
type Reader struct {
	chunkID       []byte // 0...3
	chunkSize     uint32 // 4...7
	format        []byte // 8...11
	subchunk1ID   []byte // 12...15
	subchunk1Size uint32 // 16...19
	audioFormat   uint16 // 20...21
	numChannels   uint16 // 22...23
	sampleRate    uint32 // 24...27
	byteRate      uint32 // 28...31
	blockAlign    uint16 // 32...33
	bitsPerSample uint16 // 34...35
	subchunk2ID   []byte // 36...39
	subchunk2Size uint32 // 40...43
	reader        io.Reader
}

func (r *Reader) Read(data []byte) (n int, err error) {
	return r.Read(data)
}

// GetChannels ...
func (r *Reader) GetChannels() uint16 {
	return r.numChannels
}

// GetRate ...
func (r *Reader) GetRate() uint32 {
	return r.sampleRate
}

// NewReader wav file
func NewReader(data []byte) (r *Reader, err error) {
	if len(data) < 44 {
		err = ErrSmallSize
		return
	}

	if bytes.Compare(data[0:4], tokenRiff) != 0 &&
		bytes.Compare(data[8:12], tokenWAVE) != 0 &&
		bytes.Compare(data[12:16], tokenFmt) != 0 &&
		bytes.Compare(data[36:40], tokenList) != 0 {
		err = ErrFormat
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
