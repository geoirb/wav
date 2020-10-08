package wav

var (
	tokenRiff = []byte{'R', 'I', 'F', 'F'}
	tokenWAVE = []byte{'W', 'A', 'V', 'E'}
	tokenFmt  = []byte{'f', 'm', 't', ' '}
	tokenList = []byte{'d', 'a', 't', 'a'}
)

// bitsPerSample
const (
	FormatS16LE    = 16
	SizePCM        = 16
	AudioFormatPCM = 1
)