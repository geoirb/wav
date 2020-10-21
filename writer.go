package wav

import (
	"encoding/binary"
	"os"
)

// Writer wav file
type Writer struct {
	data []byte
	size uint32
	file *os.File
}

// GetData get audio byte array
func (w *Writer) GetData() []byte {
	return w.data
}

// Write audio signal
func (w *Writer) Write(data []byte) (n int, err error) {
	n = len(data)
	w.size += uint32(n)
	w.data = append(w.data, data...)
	return
}

// Close audio file
func (w *Writer) Close() (err error) {
	binary.LittleEndian.PutUint32(w.data[4:], w.size+36)
	binary.LittleEndian.PutUint32(w.data[40:], w.size)

	if _, err = w.file.Write(w.data); err == nil {
		err = w.file.Close()
	}
	return
}

// NewWriter for recording audio data in wav format
func NewWriter(file *os.File, channels uint16, rate, format uint32) (w *Writer) {
	w = &Writer{
		file: file,
		data: make([]byte, 44),
	}
	copy(w.data[0:], tokenRiff)
	copy(w.data[8:], tokenWAVE)
	copy(w.data[12:], tokenFmt)

	binary.LittleEndian.PutUint32(w.data[16:], sizePCM)
	binary.LittleEndian.PutUint16(w.data[20:], audioFormatPCM)
	binary.LittleEndian.PutUint16(w.data[22:], channels)
	binary.LittleEndian.PutUint32(w.data[24:], rate)

	binary.LittleEndian.PutUint32(w.data[28:], uint32(channels)*rate*uint32(format)/8)
	binary.LittleEndian.PutUint16(w.data[32:], uint16(format)/8*channels)
	binary.LittleEndian.PutUint32(w.data[34:], format)
	copy(w.data[36:], tokenList)
	return
}
