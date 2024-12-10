package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2024/internal/aocinput"
	"github.com/unixlab/AoC2024/internal/dayXX"
)

// dayXXCmd represents the dayXX command
var dayXXCmd = &cobra.Command{
	Use: "dayXX",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aocinput.Read("", cmd.Use, example)
		fmt.Printf("dayXX part 1 => %d\n", dayXX.RunPart1(input))
		fmt.Printf("dayXX part 2 => %d\n", dayXX.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(dayXXCmd)
}
