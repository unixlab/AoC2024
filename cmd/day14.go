package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day14"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use: "day14",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day14 part 1 => %d\n", day14.RunPart1(input))
		fmt.Printf("day14 part 2 => %d\n", day14.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)
}
