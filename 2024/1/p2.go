package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func p2() {
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

	l2or := make(map[int]int, 10000)
	for _, num := range list2 {
		l2or[num] += 1
	}

	simScore := 0
	for _, num := range list1 {
		simScore += num * l2or[num]
	}

	fmt.Printf("The total similarity score is %d\n", simScore)
}
