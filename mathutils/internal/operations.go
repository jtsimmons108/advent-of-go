package internal

func Abs64(n int64) int64 {
	if n >= 0 {
		return n
	}
	return -n
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func Gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
