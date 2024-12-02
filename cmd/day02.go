package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aoeinput"
	"github.com/unixlab/AoC2024/internal/day02"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use: "day02",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day02 part 1 => %d\n", day02.RunPart1(input))
		fmt.Printf("day02 part 2 => %d\n", day02.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}
