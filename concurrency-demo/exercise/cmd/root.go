package cmd

import (
	"concurrency-demo/exercise/concat"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// Used for flags.
	targetDir string
	output    string
	workers   int8

	rootCmd = &cobra.Command{
		Use:   "concat",
		Short: "Concatenates file contents",
		Long:  `Concatenates file contents`,
		Run: func(cmd *cobra.Command, args []string) {
			err := concat.Merge(targetDir, output, workers)
			if err != nil {
				errMsg := fmt.Errorf("Error %v", err)
				fmt.Println(errMsg.Error())
				os.Exit(1)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetDir, "dir", "d", ".", "directory containing files to merge")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "output.txt", "output file name")
	rootCmd.PersistentFlags().Int8VarP(&workers, "workers", "w", 5, "number of workers to use")
}
