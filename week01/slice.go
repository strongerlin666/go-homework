package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

func DeleteAt(s []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return nil, ErrIndexOutOfRange
	}
	res := s[:idx]
	res = append(res, s[idx+1:]...)
	return res, nil
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1, _ := DeleteAt(s, 4)
	fmt.Println(s1)
}
