package main

import (
	"fmt"

	"github.com/meleu/gordm"
)

func main() {
	// using max = 32767 to mimic the behavior of bash's $RANDOM
	fmt.Println(gordm.GenerateRandomInt(0, 32767))
}
