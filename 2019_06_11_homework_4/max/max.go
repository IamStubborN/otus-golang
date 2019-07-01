package max

import (
	"fmt"
)

func getComparator(list []interface{}) func(i, j int) bool {
	return func(i, j int) bool {
		switch list[len(list)-1].(type) {
		case string:
			return len(list[i].(string)) > len(list[j].(string))
		case int:
			return list[i].(int) > list[j].(int)
		case float64:
			return list[i].(float64) > list[j].(float64)
		}
		return false
	}
}

func findMax(sl []interface{}, less func(i, j int) bool) (interface{}, error) {

	if len(sl) == 0 {
		return nil, fmt.Errorf("slice is empty %p", sl)
	}

	maxIndex := 0
	for idx := range sl {
		if less(idx, maxIndex) {
			maxIndex = idx
		}
	}

	return sl[maxIndex], nil
}
