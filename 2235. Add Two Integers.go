package lc

func sumNaive(num1 int, num2 int) int {
	return num1 + num2
}

func sum(num1 int, num2 int) int {
	for num2 != 0 {
		carry := num1 & num2
		num1 = num1 ^ num2
		num2 = carry << 1
	}
	return num1
}
