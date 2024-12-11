package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day11"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use: "day11",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day11 part 1 => %d\n", day11.RunPart1(input))
		fmt.Printf("day11 part 2 => %d\n", day11.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
