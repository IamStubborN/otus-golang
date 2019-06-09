package shortener

import (
	"testing"
)

func TestLinksFormatter_Shorten(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		lf   *LinksFormatter
		args args
		want string
	}{
		{
			name:"Test",
			lf: NewLinksFormatter(),
			args: args { 
				url: "https://otus.ru/lessons/razrabotchik-golang/?int_source=courses_catalog&int_term=programming",
			},
			want: "https://otus.ru/rfBd5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lf.Shorten(tt.args.url); got != tt.want {
				t.Errorf("LinksFormatter.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinksFormatter_Resolve(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		lf   *LinksFormatter
		args args
		want string
	}{
		{
			name:"Test",
			lf: NewLinksFormatter(),
			args: args { 
				url: "https://otus.ru/6ti2S",
			},
			want: "https://otus.ru/lessons/razrabotchik-golang/?int_source=courses_catalog&int_term=programming",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lf.Shorten("https://otus.ru/lessons/razrabotchik-golang/?int_source=courses_catalog&int_term=programming")
			if got := tt.lf.Resolve(tt.args.url); got != tt.want {
				t.Errorf("LinksFormatter.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
