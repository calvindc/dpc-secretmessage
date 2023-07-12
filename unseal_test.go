package dpc_secretmessage

import (
	"io"
	"reflect"
	"testing"
)

func TestNewUnseal(t *testing.T) {
	type args struct {
		reader io.Reader
		nonce  *[24]byte
		secret *[32]byte
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
			if got := NewUnseal(tt.args.reader, tt.args.nonce, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnseal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnseal_ReadMessage(t *testing.T) {
	tests := []struct {
		name    string
		u       *Unseal
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.ReadMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("Unseal.ReadMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unseal.ReadMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
