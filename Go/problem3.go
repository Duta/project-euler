package main

import (
	"fmt"
)

func main() {
	var large int64 = 600851475143
	var largestFactor int64 = 0
	var i int64 = 2
	for ; i*i < large; i++ {
		if Isprime(i) && large%i == 0 {
			largestFactor = i
		}
	}
	fmt.Println(largestFactor)
}

func Isprime(n int64) bool {
	var i int64 = 2
	for ; i*i < n; i++ {
		if n%int64(i) == 0 {
			return false
		}
	}
	return true
}
