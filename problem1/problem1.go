package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here

}

func main() {
	fmt.Println(PrimeNumber(1000000007))  // true
	fmt.Println(PrimeNumber(1500450271))  // true
	fmt.Println(PrimeNumber(1000000000))  // false
	fmt.Println(PrimeNumber(10000000019)) // true
	fmt.Println(PrimeNumber(10000000033)) // true
}
