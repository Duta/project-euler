package main

import "fmt"

func main() {
	sum := 2
	prev1, prev2 := 1, 2
	for {
		next := prev1 + prev2
		if next > 4000000 {
			break
		}
		if next%2 == 0 {
			sum += next
		}
		prev1 = prev2
		prev2 = next
	}
	fmt.Println(sum)
}
