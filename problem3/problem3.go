package main

import "fmt"

func JoinArrayRemoveDuplicate(arrayA, arrayB []string) []string {
	// your code here

}

func main() {
	// Test cases
	fmt.Println(JoinArrayRemoveDuplicate([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}))
	// ["apel", "anggur", "lemon", "leci", "nanas"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}))
	// ["samsung", "apple", "sony", "xiaomi"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"football", "basketball"}, []string{"basketball", "football"}))
	// [football basketball]
}
