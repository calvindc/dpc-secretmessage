package dpc_secretmessage

import (
	"errors"
	"io"
	"sync"

	"encoding/binary"

	"golang.org/x/crypto/nacl/secretbox"
)

// Seale encrypts everythig that is written to it, each message in a seale stream has a header and body
type Seale struct {
	locker sync.Mutex
	nonce  *[StreamNonceLength]byte
	secret *[SecretKeyLength]byte
	writer io.Writer
}

// goodbye send a secret seale contain 18 bytes of zero, which is essentially a header with both fields zeroed out.
var goodbye [GoodbyeHeaderLength]byte

// NewSeal returns a Seale that writes encrypted messages to writer.
func NewSeal(nonce *[StreamNonceLength]byte, secret *[SecretKeyLength]byte, writer io.Writer) *Seale {
	return &Seale{
		nonce:  nonce,
		secret: secret,
		writer: writer,
	}
}

// WriteMessage sending a message involves encrypting the body of the message and preparing a header for it.
// Two secret boxes are used; one to protect the header and another to protect the body.
// Seal appends an encrypted and authenticated copy of message to out, which
// must not overlap message. The key and nonce pair must be unique for each
// distinct message and the output will be Overhead bytes longer than message.
func (s *Seale) WriteMessage(msg []byte) (err error) {
	if len(msg) > MaxMessageSegmentSize {
		err = errors.New("message length exceeds maximum segment size")
		return
	}
	s.locker.Lock()
	defer s.locker.Unlock()

	headerNonce := *s.nonce
	applyNonce(s.nonce)
	bodyNonce := *s.nonce
	applyNonce(s.nonce)

	bodySeale := secretbox.Seal(nil, msg, &bodyNonce, s.secret)
	bodyMAC := bodySeale[:secretbox.Overhead]
	body := bodySeale[secretbox.Overhead:]

	header := make([]byte, LengthSize+secretbox.Overhead)
	binary.BigEndian.PutUint16(header[:2], uint16(len(msg)))
	copy(header[2:], bodyMAC)
	headerSeale := secretbox.Seal(nil, header, &headerNonce, s.secret)

	_, err = s.writer.Write(headerSeale)
	if err != nil {
		return
	}
	_, err = s.writer.Write(body)
	return err
}

// WriteGoodbye write the stream ends with a special “goodbye” header.
// Because the goodbye header is authenticated it allows a receiver to tell the difference between
// the connection genuinely being finished and a man-in-the-middle forcibly resetting the underlying TCP connection.
func (s *Seale) WriteGoodbye() error {
	s.locker.Lock()
	defer s.locker.Unlock()

	_, err := s.writer.Write(secretbox.Seal(nil, goodbye[:], s.nonce, s.secret))

	return err
}

// applyNonce for each message needs a unique nonce
func applyNonce(s *[StreamNonceLength]byte) *[StreamNonceLength]byte {
	var i int
	for i = len(s) - 1; i >= 0 && s[i] == 0xff; i-- {
		s[i] = 0
	}
	if i < 0 {
		return s
	}
	s[i] = s[i] + 1
	return s
}
