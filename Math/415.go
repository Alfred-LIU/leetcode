package Math

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.
*/

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := byte(0)
	i := len(num1) - 1
	j := len(num2) - 1

	for i >= 0 || j >= 0 {
		sum := byte(0)
		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}
		sum += carry
		carry = sum / 10
		n := sum % 10
		res = string(n+'0') + res
		i--
		j--
	}
	if carry > 0 {
		res = string(carry+'0') + res
	}
	return res
}
