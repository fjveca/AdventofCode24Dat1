package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var values1, values2, distance []int
	var TotalDistance, substraction int
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("no valid input file found")
	}

	sets := strings.Split(string(content), "\r\n")
	for _, set := range sets {
		items := strings.Split(set, "   ")
		buffer, _ := strconv.Atoi(items[0])
		values1 = append(values1, buffer)
		buffer, _ = strconv.Atoi(items[1])
		values2 = append(values2, buffer)
	}
	sort.Slice(values1, func(i, j int) bool {
		return values1[i] < values1[j]
	})
	sort.Slice(values2, func(i, j int) bool {
		return values2[i] < values2[j]
	})

	for i := 0; i < len(values1); i++ {
		//fmt.Printf("total distance on point %v: %v ", i, TotalDistance)
		if values1[i] < values2[i] {
			//distance = append(distance, values2[i]-values1[i])
			substraction = values2[i] - values1[i]
			//println("value 2 is higher")

		} else {
			//distance = append(distance, values1[i]-values2[i])
			substraction = values1[i] - values2[i]
			//println("value 1 is higher")
		}
		distance = append(distance, substraction)
		TotalDistance += substraction
	}
	println("total distance between lists:", TotalDistance)
	verify := 0
	for _, separation := range distance {
		verify += separation
	}
	fmt.Println(verify, "=", TotalDistance)
}
