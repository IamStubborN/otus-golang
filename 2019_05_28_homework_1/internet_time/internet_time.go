package internet_time

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTimeFromServer(URL string) (string, error) {
	t1, err := ntp.Time(URL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Time from server %q is %02d-%02d-%d %02d:%02d:%02d\n",
		URL, t1.Day(), t1.Month(), t1.Year(),
		t1.Hour(), t1.Minute(), t1.Second()), nil
}

func GetTimeFromPC() (string, error) {
	t2 := time.Now()
	name, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Time from PC %q is %02d-%02d-%d %02d:%02d:%02d\n",
		name, t2.Day(), t2.Month(), t2.Year(),
		t2.Hour(), t2.Minute(), t2.Second()), nil
}
