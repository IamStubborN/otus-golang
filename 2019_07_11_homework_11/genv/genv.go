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
		fp := filepath.Join(path, fi.Name())
		if fi.IsDir() {
			if err := AddEnvsFromFolder(fp); err != nil {
				log.Fatal(err)
			}
		} else {
			value, err := ioutil.ReadFile(fp)
			if err != nil {
				return err
			}
			if err := os.Setenv(fi.Name(), string(value)); err != nil {
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
