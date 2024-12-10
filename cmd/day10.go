package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day10"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use: "day10",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day10 part 1 => %d\n", day10.RunPart1(input))
		fmt.Printf("day10 part 2 => %d\n", day10.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
