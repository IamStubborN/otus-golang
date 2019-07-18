package genv

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func generateEnvFilesInTempFolder() (string, error) {
	tmpDir, err := ioutil.TempDir("", "genv")
	if err != nil {
		return "", err
	}

	one, err := ioutil.TempFile(tmpDir, "genv-")
	if err != nil {
		log.Fatal(err)
	}
	defer one.Close()
	one.WriteString(one.Name())

	two, err := ioutil.TempFile(tmpDir, "genv-")
	if err != nil {
		log.Fatal(err)
	}
	defer two.Close()
	two.WriteString(two.Name())

	anotherDir, err := ioutil.TempDir(tmpDir, "genv-AnotherDir")
	if err != nil {
		return "", err
	}

	oneAnother, err := ioutil.TempFile(anotherDir, "genv-AnotherDir")
	if err != nil {
		log.Fatal(err)
	}
	defer oneAnother.Close()
	oneAnother.WriteString(oneAnother.Name())

	twoAnother, err := ioutil.TempFile(anotherDir, "genv-AnotherDir")
	if err != nil {
		log.Fatal(err)
	}
	defer twoAnother.Close()

	twoAnother.WriteString(twoAnother.Name())

	return tmpDir, nil
}

func TestAddEnvsFromFolder(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test add envs from folder 1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := generateEnvFilesInTempFolder()
			if err != nil {
				log.Fatal(err)
			}
			tt.args.path = tmpDir
			if err := AddEnvsFromFolder(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("AddEnvsFromFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			count := 0
			for _, env := range os.Environ() {
				if strings.Contains(env, "genv") {
					count++
				}
			}
			if count != 5 {
				t.Errorf("count files must be 5, not %v", count)
			}
		})
	}
}
