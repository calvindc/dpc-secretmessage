package dpc_secretmessage

import (
	"io"

	"bytes"
	"encoding/binary"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// Unseal encrypts decrypts that is read to it, each message in a seale stream has a header and body
type Unseal struct {
	nonce  *[StreamNonceLength]byte
	secret *[SecretKeyLength]byte
	reader io.Reader
	buf    [MaxMessageSegmentSize + secretbox.Overhead]byte
}

// NewUnboxer wraps the passed Reader into an Unboxer.
func NewUnseal(reader io.Reader, nonce *[24]byte, secret *[32]byte) *Unseal {
	return &Unseal{
		nonce:  nonce,
		secret: secret,
		reader: reader,
	}
}

// ReadMessage reads the next message from the underlying stream. If the next
// message was a 'goodbye', it returns io.EOF.
func (u *Unseal) ReadMessage() ([]byte, error) {
	headerNonce := *u.nonce
	applyNonce(u.nonce)
	bodyNonce := *u.nonce
	applyNonce(u.nonce)

	// read and Unseal header
	headerSeale := u.buf[:HeaderLength]
	if _, err := io.ReadFull(u.reader, headerSeale); err != nil {
		return nil, err
	}
	headerBuf := make([]byte, 0, LengthSize+secretbox.Overhead)
	header, ok := secretbox.Open(headerBuf, headerSeale, &headerNonce, u.secret)
	if !ok {
		return nil, errors.New("invalid header seale")
	}

	// zero header indicates termination and return EOF
	if bytes.Equal(header, goodbye[:]) {
		return nil, io.EOF
	}

	// read and Unseal body
	bodyLen := binary.BigEndian.Uint16(header[:LengthSize])
	if bodyLen > MaxMessageSegmentSize {
		return nil, errors.New("message length exceeds maximum segment size")
	}
	bodySeale := u.buf[:bodyLen+secretbox.Overhead]
	if _, err := io.ReadFull(u.reader, bodySeale[secretbox.Overhead:]); err != nil {
		return nil, err
	}
	// prepend with MAC from header
	copy(bodySeale, header[LengthSize:])
	msg, ok := secretbox.Open(nil, bodySeale, &bodyNonce, u.secret)
	if !ok {
		return nil, errors.New("invalid body seale ")
	}
	return msg, nil
}
