package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day12"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use: "day12",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day12 part 1 => %d\n", day12.RunPart1(input))
		fmt.Printf("day12 part 2 => %d\n", day12.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
