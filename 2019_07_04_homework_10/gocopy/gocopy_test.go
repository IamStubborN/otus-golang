package gocopy

import (
	"log"
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		offset int64
		limit  int64
		text   string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test WriteFile 1",
			args: args{
				offset: 0,
				limit:  5,
				text:   "hello",
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "Test WriteFile 2",
			args: args{
				offset: 238,
				limit:  425,
				text:   "Swift’s String and Character types provide a fast, Unicode-compliant way to work with text in your code. The syntax for string creation and manipulation is lightweight and readable, with a string literal syntax that is similar to C. String concatenation is as simple as combining two strings with the + operator, and string mutability is managed by choosing between a constant or a variable, just like any other value in Swift. You can also use strings to insert constants, variables, literals, and expressions into longer strings, in a process known as string interpolation. This makes it easy to create custom string values for display, storage, and printing.",
			},
			want:    425,
			wantErr: false,
		},
		{
			name: "Test WriteFile 3",
			args: args{
				offset: 239,
				limit:  425,
				text:   "Swift’s String and Character types provide a fast, Unicode-compliant way to work with text in your code. The syntax for string creation and manipulation is lightweight and readable, with a string literal syntax that is similar to C. String concatenation is as simple as combining two strings with the + operator, and string mutability is managed by choosing between a constant or a variable, just like any other value in Swift. You can also use strings to insert constants, variables, literals, and expressions into longer strings, in a process known as string interpolation. This makes it easy to create custom string values for display, storage, and printing.",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			from, err := os.Create(os.TempDir() + "/from")
			if err != nil {
				log.Fatal(err)
			}
			defer from.Close()

			_, err = from.WriteString(tt.args.text)
			if err != nil {
				log.Fatal(err)
			}

			to, err := os.Create(os.TempDir() + "/to")
			if err != nil {
				log.Fatal(err)
			}
			defer to.Close()

			got, err := WriteFile(os.TempDir()+"/from", os.TempDir()+"/to", tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("WriteFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
