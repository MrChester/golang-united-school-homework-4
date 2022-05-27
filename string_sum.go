package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var res []string
	var counter int = 0
	var sum int = 0
	inputStrSlice := []rune(input)

	if input == "" {
		return "", fmt.Errorf("error: %w", errorEmptyInput)
	}

	tempSlice := normalizeInputSlice(inputStrSlice)
	tempStr := string(tempSlice)
	splitStr := strings.Split(tempStr, "+")

	for i := 0; i < len(splitStr); i++ {
		if splitStr[i] != "" {
			res = append(res, splitStr[i])
		}
	}

	for i := 0; i < len(res); i++ {
		num, e := strconv.Atoi(res[i])
		if e == nil {
			counter += 1
		} else {
			return "", fmt.Errorf("strconv syntax error: %w", e)
		}
		sum += num
	}

	if counter != 2 {
		return "", fmt.Errorf("error: %w", errorNotTwoOperands)
	}

	output = strconv.Itoa(sum)

	return output, nil
}

func normalizeInputSlice(input []rune) (output []rune) {
	for i := 0; i < len(input); i++ {
		if input[i] == 45 {
			output = append(output, 43)
		}
		if input[i] != 32 {
			output = append(output, input[i])
		}
	}
	return
}
