package main

import (
	"fmt"

	"github.com/IamStubborN/otus-golang/2019_06_28_homework_1/task_2/shortener"
)

func main() {
	lf := shortener.NewLinksFormatter()
	sh := lf.Shorten("https://otus.ru/")
	fmt.Println(sh)
	rs := lf.Resolve(sh)
	fmt.Println(rs)
}
