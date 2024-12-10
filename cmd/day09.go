package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/day09"
)

// day09Cmd represents the day09 command
var day09Cmd = &cobra.Command{
	Use: "day09",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("day09 part 1 => %d\n", day09.RunPart1(input))
		fmt.Printf("day09 part 2 => %d\n", day09.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day09Cmd)
}
