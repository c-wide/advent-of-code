package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0

	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	matches := re.FindAllString(input, -1)

	re2 := regexp.MustCompile(`\d+`)
	for _, match := range matches {
		nums := re2.FindAllString(match, -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		total += num1 * num2
	}

	fmt.Printf("The total is %d\n", total)
}
