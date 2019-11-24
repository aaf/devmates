package main

import "fmt"

// Problem:
//  Given boolean array of days at work, for example [True, False, True, True,
//  False, False, False, True] and pto (number of PTOs one can take) - where
//  boolean True means paid holiday and False means you can take a PTO.
//
//  Find the maximum length of vacation an employee can take.
//
// Input:
//
//  [False, True, False, True, False, True, False], pto=2
//
// Output: 5
//
// Why?
//   If we take PTO on indices [2] and [4],
//   then we can get the maximum length vacation (consecutive True's)
// Approach:
//   sliding window left to right (with pto number of falses)
//   keep track of max vacation length taking ptos in window
//   after processing entire array print max vacation

type (
	VacationInput struct {
		DaysAtWork []bool
		PtoGiven   int
	}
	VacationOutput struct {
		PtoIndexes  []int
		MaxVacation int
	}
)

func main() {

	inputs := []VacationInput{{
		DaysAtWork: []bool{false, true, false, true, false, true, false},
		PtoGiven:   2,
	}, {
		DaysAtWork: []bool{false},
		PtoGiven:   1,
	}, {
		DaysAtWork: []bool{false, false, false, true, false, true, false, false, false, true},
		PtoGiven:   3,
	},
	}

	for _, input := range inputs {
		fmt.Printf("Days at Work: %+v\n", input.DaysAtWork)
		fmt.Printf("PTO: %+v\n", input.PtoGiven)

		result := maxVacation(input.DaysAtWork, input.PtoGiven)
		fmt.Printf("PTO Days Index: %+v\n", result.PtoIndexes)
		fmt.Printf("Max Vacation: %v\n", result.MaxVacation)
		fmt.Println()
	}
}

func maxVacation(daysAtWork []bool, totalPto int) VacationOutput {
	var result VacationOutput
	vacation := 0

	for i := 0; i < len(daysAtWork); i++ {
		ptoIndexes := make([]int, 0)
		pto := totalPto
		for j := i; j < len(daysAtWork); j++ {
			if !daysAtWork[j] {
				pto--
				ptoIndexes = append(ptoIndexes, j)
			}
			vacation++
			if pto == 0 {
				break
			}
		}

		if vacation > result.MaxVacation {
			result.MaxVacation = vacation
			result.PtoIndexes = ptoIndexes
		}

		vacation = 0
	}
	return result
}
