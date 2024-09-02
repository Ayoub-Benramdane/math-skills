package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var numbers []float64
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		count := 0
		if (num/Average(numbers) > 100 || num/Average(numbers) < -100) && len(numbers) > 1 && 10*count < len(numbers) {
			num = Average(numbers)
			count++
		}
		numbers = append(numbers, num)
		if len(numbers) == 1 {
			fmt.Println(numbers[0]/2, numbers[0]/2+numbers[0])
		} else {
			avg := Average(numbers)
			dev := Deviation(numbers)
			med := Median(numbers)
			fmt.Printf("%d %d\n", int(math.Round((avg+med)/2-dev)), int(math.Round((avg+med)/2+dev)))
		}

	}
}

func Average(pop []float64) float64 {
	var sum float64
	for _, v := range pop {
		sum += v
	}
	return sum / float64(len(pop))
}

func Median(pop []float64) float64 {
	sort.Float64s(pop)
	l := len(pop) / 2
	if len(pop)%2 == 1 {
		return pop[l]
	}
	return (pop[l] + pop[l-1]) / 2
}

func Variance(pop []float64) float64 {
	avg := Average(pop)
	var s float64
	for _, v := range pop {
		s += (v - avg) * (v - avg)
	}
	return s / float64(len(pop))
}

func Deviation(pop []float64) float64 {
	return math.Sqrt(Variance(pop))
}
