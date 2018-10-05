package gifmeta

var (
	header          = []byte("GIF")
	extension       = byte(0x21)
	extensionToken  = []byte{extension}
	trailer         = byte(0x3B)
	imageDescriptor = byte(0x2C)
	version87a      = []byte("87a")
	version89a      = []byte("89a")
)

var (
	discard = make([]byte, 2048)
)
