package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aoeinput"
	"github.com/unixlab/AoC2024/internal/day07"
)

// day07Cmd represents the day07 command
var day07Cmd = &cobra.Command{
	Use: "day07",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day07 part 1 => %d\n", day07.RunPart1(input))
		fmt.Printf("day07 part 2 => %d\n", day07.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day07Cmd)
}
