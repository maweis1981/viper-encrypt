package main

import (
	"fmt"
	"testing"
)

func Test_encryptString(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test encryptString",
			args: args{
				plaintext: "mysecret",
				key:       "retencryptionkey",
			},
			want:    "CZlrNPtPyTVOrKWnEQEyTlWO9ibPvY",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptString(tt.args.plaintext, tt.args.key)
			beforeEncrypt, err := decrypt(got, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if beforeEncrypt != tt.args.plaintext {
				t.Errorf("encryptString() got = %v, want %v", got, tt.want)
			}
			fmt.Println(decrypt(got, tt.args.key))
			fmt.Println("------------------")
		})
	}
}
