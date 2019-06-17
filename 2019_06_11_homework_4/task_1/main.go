package main

import (
	"fmt"
	"sort"
)

func main() {

	list1 := []interface{}{64, 84, 58, 95, 92, 61, 15, 10}
	list2 := []interface{}{"two", "three", "one"}
	list3 := []interface{}{125.8, 250.2, 14.2}

	result1 := findMax(list1)
	fmt.Println(result1)

	result2 := findMax(list2)
	fmt.Println(result2)

	result3 := findMax(list3)
	fmt.Println(result3)
}

func findMax(sl []interface{}) interface{} {

	getComparator := func(list []interface{}) func(i, j int) bool {
		return func(i, j int) bool {
			switch list[0].(type) {
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

	sort.Slice(sl, getComparator(sl))
	return sl[0]
}
