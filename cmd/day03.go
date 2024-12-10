package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day03"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use: "day03",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day := cmd.Use
		if example {
			day += "p1"
		}
		input := aocinput.Read("", day, example)
		fmt.Printf("day03 part 1 => %d\n", day03.RunPart1(input))
		if example {
			day = cmd.Use + "p2"
		}
		input = aocinput.Read("", day, example)
		fmt.Printf("day03 part 2 => %d\n", day03.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
