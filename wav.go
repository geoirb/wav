package wav

var (
	tokenRiff = []byte{'R', 'I', 'F', 'F'}
	tokenWAVE = []byte{'W', 'A', 'V', 'E'}
	tokenFmt  = []byte{'f', 'm', 't', ' '}
	tokenList = []byte{'d', 'a', 't', 'a'}
)

const (
	// S16 16 bits per a sample
	S16 = 16

	sizePCM        = 16
	audioFormatPCM = 1
)
