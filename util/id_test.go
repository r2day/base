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
