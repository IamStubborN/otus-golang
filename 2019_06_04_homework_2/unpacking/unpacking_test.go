package unpacking

import "testing"

func TestUnpackingString(t *testing.T) {
	type args struct {
		test string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"a4bc2d5e"},
			want: "aaaabccddddde",
		},
		{
			name: "Test 2",
			args: args{"abcd"},
			want: "abcd",
		},
		{
			name: "Test 3",
			args: args{"45"},
			want: "",
		},
		{
			name: "Test 4",
			args: args{`qwe\4\5`},
			want: "qwe45",
		},
		{
			name: "Test 5",
			args: args{`qwe\45`},
			want: "qwe44444",
		},
		{
			name: "Test 6",
			args: args{`qwe\\5`},
			want: `qwe\\\\\`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnpackingString(tt.args.test); got != tt.want {
				t.Errorf("UnpackingString() = %v, want %v", got, tt.want)
			}
		})
	}
}
