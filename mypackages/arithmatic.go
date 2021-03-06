package arithmatic

// Checks if a number is prime or not
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func IsOddAndPrime(num int) bool {
	if num%2 != 0 && IsPrime(num) {
		return true
	} else {
		return false
	}
}
