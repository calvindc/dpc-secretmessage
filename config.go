package dpc_secretmessage

import "golang.org/x/crypto/poly1305"

const (
	// SecretKeyLength, see golang.org/x/crypto/nacl
	SecretKeyLength = 32

	// NonceLength, see golang.org/x/crypto/nacl
	StreamNonceLength = 24

	LengthSize = 2

	// HeaderLength defines the length of the header packet before the body
	HeaderLength = LengthSize + poly1305.TagSize + poly1305.TagSize

	// MaxMessageSize defines the maximum body size for seal packets
	MaxMessageSegmentSize = 8 * 1024

	// GoodbyeHeaderLength defines the stream ends(length) with a special “goodbye” header.
	GoodbyeHeaderLength = LengthSize + poly1305.TagSize
)
