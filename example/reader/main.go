package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

// WAVReader ...
type WAVReader struct {
	Channels      uint16 // 22...23
	Rate          uint32 // 24...27
	ByteRate      uint32 // 28...31
	BlockAlign    uint16 // 32...33
	BitsPerSample uint16 // 34...35

}

func NewWAV(data []byte) *WAVReader {
	return &WAVReader{
		Channels:      binary.LittleEndian.Uint16(data[22:24]),
		Rate:          binary.LittleEndian.Uint32(data[24:28]),
		ByteRate:      binary.LittleEndian.Uint32(data[28:32]),
		BlockAlign:    binary.LittleEndian.Uint16(data[32:34]),
		BitsPerSample: binary.LittleEndian.Uint16(data[34:36]),
	}
}

func main() {
	data, _ := ioutil.ReadFile("/home/geo/go/src/github.com/geoirb/sound-ethernet-streaming/example/play-file/test.wav")

	fmt.Println(NewWAV(data))

}
