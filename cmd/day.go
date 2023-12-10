package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Run the solution for a specific day",
	Long: `This command will run the solution for a specific day of the
	Advent of Code 2023 challenge.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day called")

		if len(args) == 0 {
			fmt.Println("Please provide a day number")
			return
		}

		dayNumber := args[0]
		fmt.Println("Day number:", dayNumber)
		switch dayNumber {
		case "1":
			// check for inptut.txt in command line args
			var inputFile string

			if len(args) != 2 {
				// assume input.txt is right here.
				_, err := os.Stat("input.txt")
				if err != nil {
					fmt.Println("Please provide a path to the input file")
					return
				}
				inputFile = "input.txt"
			} else {
				inputFile = args[1]
			}
			fmt.Println("Input file:", inputFile)
			// run day1Part1 with input.txt
		}
	},
}
