package util

import (
	"encoding/json"
	"fmt"
)

func json2Slice(data string) [][]int {
	var slice [][]int
	err := json.Unmarshal([]byte(data), &slice)

	if err != nil {
		fmt.Println(slice)
	}
	return slice
}
