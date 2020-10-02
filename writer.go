package wav

import (
	"encoding/binary"
	"os"
)

// Writer wav file
type Writer struct {
	wav  []byte
	size uint32
	file *os.File
}

// Write audio signal
func (w *Writer) Write(data []byte) (n int, err error) {
	n = len(data)
	w.size += uint32(n)
	w.wav = append(w.wav, data...)
	return
}

// Close audio file
func (w *Writer) Close() (err error) {
	binary.LittleEndian.PutUint32(w.wav[4:], w.size+36)
	binary.LittleEndian.PutUint32(w.wav[40:], w.size)

	w.file.Write(w.wav)
	w.file.Close()
	return
}

// NewWriter audio file
func NewWriter(file *os.File, channels uint16, rate, format uint32) (w *Writer) {
	w = &Writer{
		file: file,
		wav:  make([]byte, 44, 44),
	}
	copy(w.wav[0:], tokenRiff)
	copy(w.wav[8:], tokenWAVE)
	copy(w.wav[12:], tokenFmt)

	binary.LittleEndian.PutUint32(w.wav[16:], 16)
	binary.LittleEndian.PutUint16(w.wav[20:], 1)
	binary.LittleEndian.PutUint16(w.wav[22:], channels)
	binary.LittleEndian.PutUint32(w.wav[24:], rate)

	binary.LittleEndian.PutUint32(w.wav[28:], uint32(channels)*rate*uint32(format)/8)
	binary.LittleEndian.PutUint16(w.wav[32:], uint16(format)/8*channels)
	binary.LittleEndian.PutUint32(w.wav[34:], format)
	copy(w.wav[36:], tokenList)
	return
}
