package wav

var (
	tokenRiff = []byte{'R', 'I', 'F', 'F'}
	tokenWAVE = []byte{'W', 'A', 'V', 'E'}
	tokenFmt  = []byte{'f', 'm', 't', ' '}
	tokenList = []byte{'d', 'a', 't', 'a'}
)

const (
	// S16LE Signed 16 bit Little Endian
	S16LE = 16
	// S24LE Signed 24 bit Little Endian
	S24LE = 24
	// S32LE Signed 32 bit Little Endian
	S32LE = 32

	sizePCM        = 16
	audioFormatPCM = 1
)
