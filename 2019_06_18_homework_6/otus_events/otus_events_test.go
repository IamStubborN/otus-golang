package otus_events

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogOtusEvent(t *testing.T) {
	type args struct {
		e OtusEvent
	}
	tests := []struct {
		name  string
		args  args
		wantW bool
	}{
		{
			name: "Test Events 1",
			args: args{
				e: &HwSubmitted{Id: 3456, Code: "200", Comment: "please take a look at my homework"},
			},
			wantW: true,
		},
		{
			name: "Test Events 2",
			args: args{
				e: &HwAccepted{Id: 3456, Grade: 4},
			},
			wantW: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			LogOtusEvent(tt.args.e, w)
			if gotW := strings.Contains(w.String(), "submitted"); gotW != tt.wantW {
				t.Errorf("LogOtusEvent() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
