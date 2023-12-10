package day

import (
	"bufio"
	"fmt"
	"os"
)

func RunDay1Part1(input string) int {
	summation := 0
	// load in the input file
	readFile, err := os.Open(input)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return -1
	}
	defer readFile.Close()
	// split the input file into lines
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// search the line for the first number
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				// convert the number to an int
				summation += int(line[i]-'0') * 10
				break
			}
		}
		// search the line for the second number
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				// convert the number to an int
				summation += int(line[i] - '0')
				break
			}
		}
	}
	// return the summation
	fmt.Println("Summation:", summation)
	return summation
}

func RunDay1Part2(input string) int {
	fmt.Println("todo: implement day1 part2")
	return 0
}
