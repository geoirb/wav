# wav

Simple package for work with wav

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
w := wav.NewWriter(file, 2, 44100, wav.FormatS16LE)
for{
    // read audio data
    // ...
    w.Write(w)
    // ...
}
w.Close()
```