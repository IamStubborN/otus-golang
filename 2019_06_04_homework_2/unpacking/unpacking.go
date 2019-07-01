package unpacking

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const BackslashIndex = 92

func main() {
	fmt.Println(UnpackingString(`qwe\\5`))
}

func UnpackingString(test string) (result string) {
	sl := []rune(test)
	length := len(sl) - 1
	for idx, r := range sl {
		if unicode.IsLetter(r) {
			if idx < length && unicode.IsDigit(sl[idx+1]) {
				sliceNumbers := getDigitSlice(sl, idx+1)
				numberDigit, err := strconv.Atoi(string(sliceNumbers))
				if err != nil {
					log.Fatal(err)
				}
				result += strings.Repeat(string(r), numberDigit)
			} else {
				result += string(r)
			}
		} else if r == BackslashIndex {
			if sl[idx+1] == BackslashIndex {
				result += string(sl[idx+1])
			} else if unicode.IsDigit(sl[idx+1]) {
				sliceNumbers := getDigitSlice(sl, idx+1)
				if len(sliceNumbers) > 1 {
					numberDigit, err := strconv.Atoi(string(sliceNumbers[1:]))
					if err != nil {
						log.Fatal(err)
					}
					result += strings.Repeat(string(sliceNumbers[0]), numberDigit)
				} else if r == BackslashIndex &&
					unicode.IsDigit(sl[idx+1]) &&
					sl[idx-1] == BackslashIndex {
					numberDigit, err := strconv.Atoi(string(sliceNumbers))
					if err != nil {
						log.Fatal(err)
					}
					result += strings.Repeat(`\`, numberDigit-1)
				} else {
					result += string(sl[idx+1])
				}
			}
		}
	}
	return result
}

func getDigitSlice(sl []rune, start int) []rune {
	digitStartIdx := start
	for len(sl) > digitStartIdx && unicode.IsDigit(sl[digitStartIdx]) {
		digitStartIdx++
	}
	return sl[start:digitStartIdx]
}
