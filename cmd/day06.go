package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day06"
)

// day06Cmd represents the day06 command
var day06Cmd = &cobra.Command{
	Use: "day06",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day06 part 1 => %d\n", day06.RunPart1(input))
		fmt.Printf("day06 part 2 => %d\n", day06.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day06Cmd)
}
