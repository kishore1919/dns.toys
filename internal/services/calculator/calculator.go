package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Evaluate parses and evaluates a simple arithmetic expression with operator precedence.
func Evaluate(expr string) (float64, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	if expr == "" {
		return 0, errors.New("empty expression")
	}

	var nums []float64
	var ops []rune
	var numStr strings.Builder

	for i, c := range expr {
		if unicode.IsDigit(c) || c == '.' {
			numStr.WriteRune(c)
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			if numStr.Len() == 0 {
				return 0, fmt.Errorf("invalid syntax at position %d", i)
			}
			n, err := strconv.ParseFloat(numStr.String(), 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %v", err)
			}
			nums = append(nums, n)
			ops = append(ops, c)
			numStr.Reset()
		} else {
			return 0, fmt.Errorf("invalid character: %c", c)
		}
	}
	if numStr.Len() == 0 {
		return 0, errors.New("expression ends with operator")
	}
	n, err := strconv.ParseFloat(numStr.String(), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}
	nums = append(nums, n)

	// First pass: handle multiplication and division
	newNums := []float64{nums[0]}
	newOps := []rune{}

	for i, op := range ops {
		if op == '*' || op == '/' {
			lastNum := newNums[len(newNums)-1]
			nextNum := nums[i+1]
			if op == '*' {
				newNums[len(newNums)-1] = lastNum * nextNum
			} else {
				if nextNum == 0 {
					return 0, errors.New("division by zero")
				}
				newNums[len(newNums)-1] = lastNum / nextNum
			}
		} else {
			newNums = append(newNums, nums[i+1])
			newOps = append(newOps, op)
		}
	}

	// Second pass: handle addition and subtraction
	result := newNums[0]
	for i, op := range newOps {
		switch op {
		case '+':
			result += newNums[i+1]
		case '-':
			result -= newNums[i+1]
		default:
			return 0, fmt.Errorf("unsupported operator: %c", op)
		}
	}
	return result, nil
}
