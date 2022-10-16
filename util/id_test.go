package util

import "testing"

func TestSmsCode(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"test",
			"R-157807",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmsCode(); got != tt.want {
				t.Errorf("SmsCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getId(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				prefix: "R-",
			},
			"R-xx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getId(tt.args.prefix); got != tt.want {
				t.Errorf("getId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToToken(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				"hello",
			},
			"5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToToken(tt.args.k); got != tt.want {
				t.Errorf("ConvertToToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
