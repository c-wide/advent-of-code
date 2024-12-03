package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func p2() {
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

	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)|do(?:n't)?`)
	matches := re.FindAllString(input, -1)

	state := 1

	re2 := regexp.MustCompile(`\d+`)
	for _, match := range matches {
		switch match {
		case "do":
			state = 1
		case "don't":
			state = 0
		default:
			if state == 1 {
				nums := re2.FindAllString(match, -1)
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				total += num1 * num2
			}

		}
	}

	fmt.Printf("The total is %d\n", total)
}
