package main

import (
	"fmt"
)

func MunculSekali(angka string) []int {
	// your code here

}

func main() {
	fmt.Println(MunculSekali("1234123"))    // [4]
	fmt.Println(MunculSekali("76523752"))   // [6, 3]
	fmt.Println(MunculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // []
	fmt.Println(MunculSekali("0872504"))    // [8 7 2 5 4]
}
