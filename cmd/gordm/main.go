package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/meleu/gordm"
)

const (
	defaultMin = 0
	defaultMax = 32767 // using max = 32767 to mimic the behavior of bash's $RANDOM
)

type parsedArgs struct {
	min int
	max int
}

func parseArgs() parsedArgs {
	min := flag.Int("min", defaultMin, "Minimum number to be generated")
	max := flag.Int("max", defaultMax, "Maximum number to be generated")
	flag.Parse()

	validateMinMax(*min, *max)

	return parsedArgs{
		min: *min,
		max: *max,
	}
}

func main() {
	args := parseArgs()
	fmt.Println(gordm.GenerateRandomInt(args.min, args.max))
}

func validateMinMax(min, max int) {
	if min > max {
		fmt.Fprintf(
			os.Stderr,
			"invalid input: 'min' cannot be greater than 'max' (got min=%d, max=%d)",
			min, max,
		)
		os.Exit(1)
	}
}
