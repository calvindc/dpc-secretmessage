package dpc_secretmessage

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewSeal(t *testing.T) {
	type args struct {
		nonce  *[StreamNonceLength]byte
		secret *[SecretKeyLength]byte
	}
	tests := []struct {
		name       string
		args       args
		want       *Seale
		wantWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if got := NewSeal(tt.args.nonce, tt.args.secret, writer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeal() = %v, want %v", got, tt.want)
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("NewSeal() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}

func TestSeale_WriteMessage(t *testing.T) {
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		s       *Seale
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.WriteMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Seale.WriteMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSeale_WriteGoodbye(t *testing.T) {
	tests := []struct {
		name    string
		s       *Seale
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.WriteGoodbye(); (err != nil) != tt.wantErr {
				t.Errorf("Seale.WriteGoodbye() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_applyNonce(t *testing.T) {
	type args struct {
		s *[StreamNonceLength]byte
	}
	tests := []struct {
		name string
		args args
		want *[StreamNonceLength]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyNonce(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}
