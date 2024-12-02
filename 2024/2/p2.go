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
			continue
		}

		for idx := range report {
			newReport := append([]int{}, report[:idx]...)
			newReport = append(newReport, report[idx+1:]...)
			if isValidReport(newReport) {
				safeReports++
				break
			}
		}
	}

	fmt.Printf("There are %d safe reports\n", safeReports)
}
