package funcs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func File(a string) []float64 {
	file, err := os.ReadFile(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	k := strings.Split(string(file), "\n")
	s := make([]float64, 0)
	for _, v := range k {
		if v == "" {
			continue
		}
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		s = append(s, (i))
	}
	return s
}
