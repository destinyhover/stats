/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search command",
	Long:  `The search command search for info`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Searching for")
		logger.Info(sid)

		for i, k := range data {
			if k.Filename == sid {
				fmt.Println(data[i])
				break
			}
		}
	},
}

var sid string

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&sid, "sid", "s", "", "Search Key")
	searchCmd.MarkFlagRequired("sid")
}
