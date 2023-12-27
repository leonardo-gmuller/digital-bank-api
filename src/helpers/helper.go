package helpers

import (
	"strconv"
	"strings"
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
	if len(newArray) != 11 {
		return false
	}
	first := newArray[:9]
	second := newArray[9:]
	if cpf[0] == cpf[1] && cpf[1] == cpf[2] && (cpf[2] == cpf[3]) && (cpf[3] == cpf[4]) && (cpf[4] == cpf[5]) && (cpf[5] == cpf[6]) && (cpf[6] == cpf[7]) && (cpf[7] == cpf[8]) && (cpf[8] == cpf[9]) && (cpf[9] == cpf[10]) {
		return false
	} else {

		sum1 := first[0]*10 + first[1]*9 + first[2]*8 + first[3]*7 + first[4]*6 + first[5]*5 + first[6]*4 + first[7]*3 + first[8]*2
		sum2 := first[0]*11 + first[1]*10 + first[2]*9 + first[3]*8 + first[4]*7 + first[5]*6 + first[6]*5 + first[7]*4 + first[8]*3 + second[0]*2
		rest1 := (sum1 * 10) % 11
		rest2 := (sum2 * 10) % 11
		if rest1 == 10 {
			rest1 = 0
		}
		if rest2 == 10 {
			rest2 = 0
		}
		if rest1 != second[0] && rest2 != second[1] {
			return false
		}
		return true
	}
}
