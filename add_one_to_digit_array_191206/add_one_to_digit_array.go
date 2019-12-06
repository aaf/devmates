package main

//Given a non-empty array of digits representing a non-negative integer and each
//element in the array contain a single digit.

//Add one (1) to this integer and return it as array.

//You may assume the integer does not contain any leading zero, except the
//number 0 itself.

//Note: Converting to the integer and back to the array is not allowed.

//Input: [3,2,1]
//Output: [3,2,2]

//Input:[9,9,9]
//Output: [1,0,0,0]

//Input:[5,6,9]
//Output: [5,7,0]

import (
	"fmt"
	"strconv"
)

func main() {
	input := [][]int{
		{3, 2, 1},
		{9, 9, 9},
		{5, 6, 9}}

	for _, row := range input {
		i := arrayToInt(row)
		fmt.Printf("input digit: %v\n", i)
		i++
		fmt.Printf("resulting array: %v\n", intToArray(i))
	}
}

func arrayToInt(input []int) int {
	res := 0
	mult := 1
	for i := len(input) - 1; i >= 0; i-- {
		res = res + (input[i] * mult)
		mult = mult * 10
	}
	return res
}

func intToArray(input int) []int {
	var res []int
	str := strconv.Itoa(input)
	for _, char := range str {
		res = append(res, int(char-'0'))
	}
	return res
}
