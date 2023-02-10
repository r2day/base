package time

import "testing"

func TestFormatTime(t *testing.T) {
	type args struct {
		timestamp int64
		layout    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				1676068589,
				"2006-01-02 15:04:05",
			},
			"2023-02-11 06:36:29",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTime(tt.args.timestamp, tt.args.layout); got != tt.want {
				t.Errorf("FormatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
