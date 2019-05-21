package multiply_strings

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) []int {
	arr := make([]int, len(num1)+len(num2))

	carry := 0
	carry2 := 0
	itr := 0
	for i := len(num2) - 1; i >= 0; i-- {

		k := len(arr) - itr - 1

		for _, n2 := range num1 {

			n1, _ := strconv.Atoi(string(num2[i]))
			in2, _ := strconv.Atoi(string(n2))

			fmt.Println("INDATA :", n1, in2, k)

			product := (n1 * in2) + carry
			n := product % 10
			carry = product / 10

			sum := n + arr[k] + carry2
			nn := sum % 10
			carry2 = sum / 10

			arr[k] = nn
			k--
		}

		if carry > 0 {
			if carry == 10 {
				arr[k] = 0
				arr[k-1] = 1
			}

			arr[k] = carry
		}

		itr++

	}

	return arr
}
