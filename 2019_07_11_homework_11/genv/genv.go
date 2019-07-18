package genv

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func AddEnvsFromFolder(path string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fi := range dir {
		if fi.IsDir() {
			fp := filepath.Join(path, fi.Name())
			if err := AddEnvsFromFolder(fp); err != nil {
				log.Fatal(err)
			}
		} else {
			key := fi.Name()
			value, err := ioutil.ReadFile(filepath.Join(path, fi.Name()))
			if err != nil {
				return err
			}
			if err := os.Setenv(key, string(value)); err != nil {
				return err
			}
		}
	}

	return nil
}

func ExecuteCmdWithEnv(name string) error {
	cmd := exec.Command(name)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
