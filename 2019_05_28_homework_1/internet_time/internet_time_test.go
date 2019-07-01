package internet_time

import (
	"strings"
	"testing"
)

func TestGetTimeFromServer(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Internet time 1",
			args: args{
				"0.beevik-ntp.pool.ntp.org",
			},
			want:    "0.beevik-ntp.pool.ntp.org",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimeFromServer(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimeFromServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("GetTimeFromServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTimeFromPC(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "PC time 1",
			want:    "PC",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimeFromPC()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimeFromPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("GetTimeFromServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
