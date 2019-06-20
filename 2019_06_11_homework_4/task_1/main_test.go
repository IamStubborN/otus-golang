package main

import (
	"reflect"
	"testing"
)

func Test_findMax(t *testing.T) {
	type args struct {
		sl []interface{}
	}
	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
		args    args
	}{
		{
			name:    "findMax 1. Ints.",
			wantErr: false,
			want:    95,
			args: args{
				sl: []interface{}{64, 84, 58, 95, 92, 61, 15, 10},
			},
		},
		{
			name:    "findMax 2. Floats.",
			wantErr: false,
			want:    250.2,
			args: args{
				sl: []interface{}{125.8, 250.2, 14.2},
			},
		},
		{
			name:    "findMax 3. Strings.",
			wantErr: false,
			want:    "three",
			args: args{
				sl: []interface{}{"two", "three", "one"},
			},
		},
		{
			name:    "findMax 4. Nil slice.",
			wantErr: true,
			want:    nil,
			args: args{
				sl: []interface{}{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findMax(tt.args.sl, getComparator(tt.args.sl))
			if (err != nil) != tt.wantErr {
				t.Errorf("findMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
