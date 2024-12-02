package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidReport(report []int) bool {
	state := 0

	for currIdx := 0; currIdx < len(report)-1; currIdx++ {
		currNum := report[currIdx]
		nextNum := report[currIdx+1]

		if currNum == nextNum {
			return false
		}

		if currNum > nextNum {
			diff := currNum - nextNum
			if state == 1 || diff > 3 {
				return false
			}
			state = -1
		} else {
			diff := nextNum - currNum
			if state == -1 || diff > 3 {
				return false
			}
			state = 1

		}
	}

	return true
}

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		temp := make([]int, 0, len(values))
		for _, strNum := range values {
			num, _ := strconv.Atoi(strNum)
			temp = append(temp, num)
		}
		reports = append(reports, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	safeReports := 0
	for _, report := range reports {
		if isValidReport(report) {
			safeReports++
		}
	}

	fmt.Printf("There are %d safe reports\n", safeReports)
}
