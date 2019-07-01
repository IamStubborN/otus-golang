package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IamStubborN/otus-golang/2019_06_04_homework_2/unpacking"
	"github.com/IamStubborN/otus-golang/2019_06_06_homework_3/words"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Hello. What are you want to use?\n1) Unpacking string\n2) Ten popular words\n3) Exit\n")
		scanner.Scan()
		ans := scanner.Text()
		switch ans {
		case "1":
			fmt.Printf("Insert the line you want to unpack\n")
			scanner.Scan()
			ans = scanner.Text()
			fmt.Printf("\n-----" + unpacking.UnpackingString(ans) + "-----\n\n")
		case "2":
			fmt.Printf("Insert the text you want to calc popular words\n")
			scanner.Scan()
			ans = scanner.Text()
			result, err := words.TenPopularWordsFromText(ans)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("\n-----" + strings.Join(result, " ") + "-----\n\n")
		case "3":
			fmt.Printf("Bye!\n")
			os.Exit(0)
		default:
			fmt.Printf("Bad request\n")
		}
	}
}
