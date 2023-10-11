/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nikitarudakov/tac/file"
	"github.com/nikitarudakov/tac/logger"
	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "unzip",
	Short: "Unzip file",
	Long:  `To unzip file pass file/to/path as an argument`,

	Run: func(cmd *cobra.Command, args []string) {

		logger := logger.InitLogger()

		logger.Info().Msg("Unzip command was called")

		if len(args) > 1 {
			if err := file.DecompressZipFile(args[0], args[1]); err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
