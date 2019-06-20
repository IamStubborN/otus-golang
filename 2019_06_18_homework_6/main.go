package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type (
	HwAccepted struct {
		Id    int
		Grade int
	}

	HwSubmitted struct {
		Id      int
		Code    string
		Comment string
	}

	OtusEvent interface {
		Log(io.Writer)
	}
)

func (ha *HwAccepted) Log(w io.Writer) {
	fmt.Fprintf(w, "%s accepted %d %d\n", getDate(), ha.Id, ha.Grade)
}

func (ha *HwSubmitted) Log(w io.Writer) {
	fmt.Fprintf(w, "%s submitted %d %q\n", getDate(), ha.Id, ha.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	e.Log(w)
}

func main() {

	var e1, e2 OtusEvent

	e1 = &HwSubmitted{Id: 3456, Code: "200", Comment: "please take a look at my homework"}
	e2 = &HwAccepted{Id: 3456, Grade: 4}

	LogOtusEvent(e1, os.Stdout)
	LogOtusEvent(e2, os.Stdout)
}

func getDate() string {
	return time.Now().Format("2006-01-02")
}
