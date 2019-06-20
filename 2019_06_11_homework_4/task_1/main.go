package main

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

func main() {

	list1 := []interface{}{64, 84, 58, 95, 92, 61, 15, 10}
	list2 := []interface{}{"two", "three", "one"}
	list3 := []interface{}{125.8, 250.2, 14.2}
	list4 := []interface{}{}

	result1, err := findMax(list1, getComparator(list1))
	fmt.Println(result1, err)

	result2, err := findMax(list2, getComparator(list2))
	fmt.Println(result2, err)

	result3, err := findMax(list3, getComparator(list3))
	fmt.Println(result3, err)

	result4, err := findMax(list4, getComparator(list4))
	fmt.Println(result4, err)
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
