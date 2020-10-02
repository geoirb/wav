package wav

import (
	"encoding/binary"
)

// wav file fields
var (
	chunkID               = []byte{'R', 'I', 'F', 'F'}
	format                = []byte{'W', 'A', 'V', 'E'}
	subchunk1ID           = []byte{'f', 'm', 't', ' '}
	subchunk2ID           = []byte{'L', 'I', 'S', 'T'}
	pcmSize        uint16 = 16
	pcmAudioFormat uint16 = 1

	FormatS16LE uint16 = 16
)

// WAV info
type WAV struct {
	chunkID       []byte // 0...3
	chunkSize     uint32 // 4...7
	format        []byte // 8...11
	subchunk1ID   []byte // 12...15
	subchunk1Size uint16 // 16...19
	audioFormat   uint16 // 20...21
	numChannels   uint16 // 22...23
	sampleRate    uint32 // 24...27
	byteRate      uint32 // 28...31
	blockAlign    uint16 // 32...33
	bitsPerSample uint16 // 34...35
	subchunk2ID   []byte // 36...39
	subchunk2Size uint32 // 40...43
}

// GetChannels ...
func (w *WAV) GetChannels() uint16 {
	return w.numChannels
}

// GetRate ...
func (w *WAV) GetRate() uint32 {
	return w.sampleRate
}

// SetChannels ...
func (w *WAV) SetChannels(channels uint16) {
	w.numChannels = channels
}

// SetRate ...
func (w *WAV) SetRate(rate uint32) {
	w.sampleRate = rate
}

// Set16LE ...
func (w *WAV) Set16LE() {
	w.bitsPerSample = FormatS16LE
}

// NewWAV ...
func NewWAV() *WAV {
	return &WAV{
		chunkID: chunkID,
		// chunkSize
		format:        format,
		subchunk1ID:   subchunk1ID,
		subchunk1Size: pcmSize,
		audioFormat:   pcmAudioFormat,
		subchunk2ID:   subchunk2ID,
	}
}

// ParseWAV parse info about wav file
func ParseWAV(data []byte) *WAV {
	return &WAV{
		chunkID:       data[:4],
		chunkSize:     binary.LittleEndian.Uint32(data[4:8]),
		format:        data[8:12],
		subchunk1ID:   data[12:16],
		subchunk1Size: binary.LittleEndian.Uint16(data[16:20]),
		audioFormat:   binary.LittleEndian.Uint16(data[20:22]),
		numChannels:   binary.LittleEndian.Uint16(data[22:24]),
		sampleRate:    binary.LittleEndian.Uint32(data[24:28]),
		byteRate:      binary.LittleEndian.Uint32(data[28:32]),
		blockAlign:    binary.LittleEndian.Uint16(data[32:34]),
		bitsPerSample: binary.LittleEndian.Uint16(data[34:36]),
		subchunk2ID:   data[36:40],
		subchunk2Size: binary.BigEndian.Uint32(data[40:44]),
	}
}
