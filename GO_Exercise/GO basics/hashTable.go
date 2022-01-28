// GO language program with an example of Hash Table

package main

import (
	"fmt"
)

func main() {
	kl := make(map[int]string)
	kl[1] = "UI"
	kl[2] = "Dev"
	kl[3] = "SIG"

	for i, j := range kl {
		fmt.Printf("Hash: %d Value: %s\n", i, j)
	}
}
