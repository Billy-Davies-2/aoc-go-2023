package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "This is a command line tool for Advent of Code 2023",
	Long: `I am trying to learn how to master golang more effectively,
	and I am using Advent of Code 2023 as a way to do that. This is a
	command line tool that will allow me to run the solutions for each
	day of the Advent of Code 2023 challenge.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(dayCmd)

}
