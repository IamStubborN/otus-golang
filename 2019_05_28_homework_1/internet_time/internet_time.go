package internet_time

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const HOST_URL = "0.beevik-ntp.pool.ntp.org"

func main() {
	t1, err := GetTimeFromServer(HOST_URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1)
	t2, err := GetTimeFromPC()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t2)
}

func GetTimeFromServer(URL string) (string, error) {
	t1, err := ntp.Time(HOST_URL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Time from server %q is %02d-%02d-%d %02d:%02d:%02d\n",
		HOST_URL, t1.Day(), t1.Month(), t1.Year(),
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
