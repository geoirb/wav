# wav

Simple package for work with wav
## WAV file structure

|Offset(byte)|Size(byte)|Name|Description|
|:--:|:--:|:--:|:--:
|0|4|chunkId|Contains "RIFF" characters in ASCII encoding. It is the beginning of the RIFF chain.
|4|4|chunkSize|This is the remaining chain size from that position. In other words, this is the file size minus 8, that is, the chunkId and chunkSize fields are excluded.
|8|4|format|Contains "WAVE" symbols
|12|4|subchunk1Id|Contains the characters "fmt "
|16|4|subchunk1Size|16 for PCM format. This is the remaining size of the chain from this position.
|20|2|audioFormat|Audio format, list of acceptable formats. For PCM = 1 (that is, Linear Quantization). Values other than 1 indicate some compression format.
|22|2|numChannels|Number of channels. Mono = 1, Stereo = 2, etc.
|24|4|sampleRate|Sampling frequency. 8000 Hz, 44100 Hz, etc.
|28|4|byteRate|The number of bytes transferred per second of playback.
|32|2|blockAlign|The number of bytes for one sample, including all channels.
|34|2|bitsPerSample|The number of bits in the sample. The so-called "depth" or accuracy of sound. 8 bits, 16 bits, etc.
|36|4|subchunk2Id|Contains symbols "data"
|40|4|subchunk2Size|The number of bytes in the data area.
|44|...|data|WAV audio data.

## Example

```golang
import "github.com/geoirb/wav"
```

Read .wav file

```golang
    data, _ := ioutil.ReadFile("path-file.wav")
    wav.NewReader(data)
```

Create .wav file

```golang
	file, _ := os.Create("path-file.wav")
	w := wav.NewWriter(file, 2, 44100, wav.S16LE)
	var data []byte
	for {
		// read audio data
		// ...
		w.Write(data)
		// ...
	}
	w.Close()
```

