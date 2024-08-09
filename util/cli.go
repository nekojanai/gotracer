package util

import (
	"log"
	"os"
	"strconv"
)

func ParseIntArg(position int) int {
	arg, err := strconv.Atoi(os.Args[position])
	if err != nil {
		log.Fatal("Error when parsing samples_per_pixel:", err)
	}

	return arg
}
