package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aoeinput"
	"github.com/unixlab/AoC2024/internal/day04"
)

// day04Cmd represents the day04 command
var day04Cmd = &cobra.Command{
	Use: "day04",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day04 part 1 => %d\n", day04.RunPart1(input))
		fmt.Printf("day04 part 2 => %d\n", day04.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day04Cmd)
}
