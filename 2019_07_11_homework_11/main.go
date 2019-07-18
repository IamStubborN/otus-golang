package main

import (
	"log"

	"github.com/IamStubborN/otus-golang/2019_07_11_homework_11/genv"
	"github.com/spf13/pflag"
)

var path *string
var appName *string

func init() {
	path = pflag.StringP("path", "p", "", "genv -p ./envs")
	appName = pflag.StringP("app", "a", "", "genv -a env")
	pflag.Parse()
}

func main() {
	if err := genv.AddEnvsFromFolder(*path); err != nil {
		log.Fatal(err)
	}
	if err := genv.ExecuteCmdWithEnv(*appName); err != nil {
		log.Fatal(err)
	}
}
