/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nikitarudakov/tac/zipfile"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "unzip",
	Short: "Unzip zipfile",
	Long:  `To unzip zipfile pass zipfile/to/path as an argument`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Unzip command was called")

		// Run validator of args
		isFolder := true

		if len(args) > 1 {
			if err := zipfile.Unzip(args[0], args[1], isFolder); err != nil {
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
