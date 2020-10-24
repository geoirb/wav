package wav

import (
	"io/ioutil"
	"os"
)

// ExampleReader example using wav Reader
func ExampleReader() {
	data, _ := ioutil.ReadFile("path-file.wav")
	NewReader(data)
}

// ExampleWriter example using wav Writer
func ExampleWriter() {
	file, _ := os.Create("path-file.wav")
	w := NewWriter(file, 2, 44100, S16LE)
	var data []byte
	for {
		// read audio data
		// ...
		w.Write(data)
		// ...
	}
	w.Close()
}
