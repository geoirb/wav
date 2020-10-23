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

// GetChunkID field
func (r *Reader) GetChunkID() []byte {
	return r.chunkID
}

// GetChunkSize size from that position
func (r *Reader) GetChunkSize() uint32 {
	return r.chunkSize
}

// GetFormat field
func (r *Reader) GetFormat() []byte {
	return r.format
}

// GetSubchunk1ID field
func (r *Reader) GetSubchunk1ID() []byte {
	return r.subchunk1ID
}

// GetSubchunk1Size field
func (r *Reader) GetSubchunk1Size() uint32 {
	return r.subchunk1Size
}

// GetAudioFormat audio format
func (r *Reader) GetAudioFormat() uint16 {
	return r.audioFormat
}

// GetNumChannels number of channels
func (r *Reader) GetNumChannels() uint16 {
	return r.numChannels
}

// GetSampleRate sampling frequency
func (r *Reader) GetSampleRate() uint32 {
	return r.sampleRate
}

// GetByteRate number of bytes transferred per second of playback
func (r *Reader) GetByteRate() uint32 {
	return r.byteRate
}

// GetBlockAlign number of bytes for one sample
func (r *Reader) GetBlockAlign() uint16 {
	return r.blockAlign
}

// GetBitsPerSample number of bits in the sample
func (r *Reader) GetBitsPerSample() uint16 {
	return r.bitsPerSample
}

// GetSubchunk2ID field
func (r *Reader) GetSubchunk2ID() []byte {
	return r.subchunk2ID
}

// GetSubchunk2Size number of bytes in the data area
func (r *Reader) GetSubchunk2Size() uint32 {
	return r.subchunk2Size
}

// Read audio
func (r *Reader) Read(data []byte) (n int, err error) {
	return r.reader.Read(data)
}

// NewReader parse audio data in wav format
func NewReader(data []byte) (r *Reader, err error) {
	if len(data) < 44 {
		err = ErrSmallSize
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
