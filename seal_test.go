package dpc_secretmessage

import (
	"crypto/rand"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/poly1305"
	"io"
	"net"
	"testing"
)

const MsgSize = 8 * 1024

func TestSealAndUnseal(t *testing.T) {
	for i := 0; i < 10; i++ {
		TestSeale_Unseal(t)
	}
}

func BenchmarkSealBytes(b *testing.B) {
	benchmarkSealDifferentMsgSize(b, 3096)
}

func TestSeale_Unseal(t *testing.T) {
	//var message = []byte("hello,!@#$%^&*()_+1234567890qazyhnpl,;.,/'[]olleh")
	var message [MsgSize]byte
	rand.Reader.Read(message[:])
	t.Logf("input message is \t%v \tbefore seal", message)

	var secret [SecretKeyLength]byte
	rand.Reader.Read(secret[:])
	var nonce [StreamNonceLength]byte
	rand.Reader.Read(nonce[:])
	var unsealnonce = nonce

	pipeR, pipeW := net.Pipe()
	wErrChan := make(chan error)
	cErrChan := make(chan error)
	checkWrite := checkErr(wErrChan)
	checkEndMsg := checkErr(cErrChan)
	seale := NewSeal(&nonce, &secret, pipeW)
	go func() {
		//one message stream contains body and goodbye-tag
		err := seale.WriteMessage(message[:])
		checkWrite(err)
		err = seale.WriteGoodbye()
		checkEndMsg(err)
	}()
	unseal := NewUnseal(&unsealnonce, &secret, pipeR)
	unsealedMsg, err := unseal.ReadMessage()
	if err != nil {
		t.Fatal(err)
	}
	if errW, has := <-wErrChan; has {
		t.Fatal(errW)
	}

	//check stream body
	if len(unsealedMsg) != len(message) {
		t.Error("unsealedMsg length wrong")
	}
	for i, x := range unsealedMsg {
		if x != message[i] {
			t.Errorf("check every byte failed, expect %v, but got %v", message[i], x)
		}
	}

	t.Logf("output message is \t%v \twhen do unseal", unsealedMsg)
	//check stream end tag
	unsealedEnd, err := unseal.ReadMessage()
	if err != io.EOF {
		t.Fatal(err)
	}
	if len(unsealedEnd) != 0 {
		t.Fatal("stream unseal end with a wrong tag")
	}
	if errC, has := <-cErrChan; has {
		t.Fatal(errC)
	}

}

func benchmarkSealDifferentMsgSize(b *testing.B, msgsize int64) {
	message := make([]byte, msgsize)
	out := make([]byte, poly1305.TagSize+msgsize)
	var secret [SecretKeyLength]byte
	var nonce [StreamNonceLength]byte
	b.SetBytes(msgsize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out = secretbox.Seal(nil, message, &nonce, &secret)
	}
	b.Log(out)
	b.Logf("message size=%d,after seal the stream size=%d,len(msgsize)+16 should equal len(out)", msgsize, len(out))
}

func checkErr(err0 chan<- error) func(error) {
	return func(err error) {
		if err != nil {
			err0 <- err
		} else {
			close(err0)
		}
	}
}
