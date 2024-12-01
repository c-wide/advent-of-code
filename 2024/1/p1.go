package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list1 := make([]int, 0, 10000)
	list2 := make([]int, 0, 10000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		n1, _ := strconv.Atoi(values[0])
		n2, _ := strconv.Atoi(values[1])

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i, num := range list1 {
		totalDistance += int(math.Abs(float64(num - list2[i])))
	}

	fmt.Printf("The total distance is %d\n", totalDistance)
}
