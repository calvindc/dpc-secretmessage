package dpc_secretmessage

import (
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"reflect"
	"testing"
)

func TestNewUnseal(t *testing.T) {
	type args struct {
		nonce  *[24]byte
		secret *[32]byte
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want *Unseal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnseal(tt.args.nonce, tt.args.secret, tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnseal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnseal_ReadMessage(t *testing.T) {
	type fields struct {
		nonce  *[StreamNonceLength]byte
		secret *[SecretKeyLength]byte
		reader io.Reader
		buf    [MaxMessageSegmentSize + secretbox.Overhead]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Unseal{
				nonce:  tt.fields.nonce,
				secret: tt.fields.secret,
				reader: tt.fields.reader,
				buf:    tt.fields.buf,
			}
			got, err := u.ReadMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
