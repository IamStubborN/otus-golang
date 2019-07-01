package words

import (
	"reflect"
	"sort"
	"testing"
)

func TestTenPopularWordsFromText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test words 1",
			args: args{
				text: "i will crush you. you are will smith. will i am. currently on side.",
			},
			want:    []string{"am", "are", "crush", "currently", "i", "on", "side", "smith", "will", "you"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TenPopularWordsFromText(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TenPopularWordsFromText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenPopularWordsFromText() = %v, want %v", got, tt.want)
			}
		})
	}
}
