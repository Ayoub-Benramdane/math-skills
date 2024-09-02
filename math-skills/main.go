package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil || len(file) == 0 {
		fmt.Println("Error")
		return
	}
	split_file := strings.Split(string(file), "\n")
	numbers := []int{}
	for _, c := range split_file {
		if c == "" {
			continue
		}
		nb, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, nb)
	}
	fmt.Println("Average:", int(math.Round(average(numbers))))
	fmt.Println("Median:", int(math.Round(median(numbers))))
	fmt.Println("Variance:", int(math.Round(variance(numbers))))
	fmt.Println("Standard Deviation:", int(math.Round(Deviation(numbers))))
}

func average(numbers []int) float64 {
	res := 0
	for _, c := range numbers {
		res += c
	}
	return float64(res) / float64(len(numbers))
}

func median(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers)%2 != 0 {
		return float64(numbers[len(numbers)/2])
	}
	return (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2
}

func variance(numbers []int) float64 {
	var res float64
	for _, c := range numbers {
		res += (float64(c) - average(numbers)) * (float64(c) - average(numbers))
	}
	return float64(res) / float64(len(numbers))
}

func Deviation(numbers []int) float64 {
	return math.Sqrt(variance(numbers))
}
