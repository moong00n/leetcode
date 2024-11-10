package main

import "fmt"

func BitwiseAdd(a, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}

	return a
}

func BitwiseAbs(n int) int {
	mask := n >> 31
	return (n ^ mask) - mask
}

func isOdd(n int) bool {
	return n&1 != 0
}

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

func countBits(n int) int {
	count := 0

	for n != 0 {
		count += n & 1
		n >>= 1
	}

	return count
}

func main() {

	fmt.Println(BitwiseAdd(5, 3)) // 8
	fmt.Println(BitwiseAbs(-5))   // 5
	fmt.Println(isOdd(15))        // true
	fmt.Println(isPowerOfTwo(8))  // true
	fmt.Println(countBits(7))     // 3

}
