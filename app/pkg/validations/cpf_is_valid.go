package validations

import (
	"strconv"
	"strings"
)

const (
	multiplierFirstDigit  = 10
	multiplierSecondDigit = 11
	moduloValue           = 11
	replacementValue      = 10
	cpfLength             = 11
)

func CPFIsValid(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpfArray := strings.Split(cpf, "")
	newArray := []int{}

	for _, i := range cpfArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		newArray = append(newArray, j)
	}

	if len(newArray) != cpfLength {
		return false
	}

	first := newArray[:9]
	second := newArray[9:]

	if allDigitsEqual(newArray) {
		return false
	}

	sum1 := calculateSum(first, multiplierFirstDigit)
	x := 2
	sum2 := calculateSum(first, multiplierSecondDigit) + second[0]*x
	x1 := 10
	rest1 := (sum1 * x1) % moduloValue
	rest2 := (sum2 * x1) % moduloValue

	if rest1 == replacementValue {
		rest1 = 0
	}

	if rest2 == replacementValue {
		rest2 = 0
	}

	return rest1 == second[0] && rest2 == second[1]
}

func allDigitsEqual(cpf []int) bool {
	for i := range cpf {
		if cpf[i] != cpf[0] {
			return false
		}
	}

	return true
}

func calculateSum(digits []int, factor int) int {
	sum := 0
	for _, digit := range digits {
		sum += digit * factor
		factor--
	}

	return sum
}
