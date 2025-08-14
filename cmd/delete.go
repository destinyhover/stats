/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete command",
	Long:  `A command for deliting data`,
	Run: func(cmd *cobra.Command, args []string) {
		_, ok := index[key]
		if ok {
			logger.Info("Found key:", key)
			delete(index, key)
		} else {
			s := fmt.Sprintf("%s not found!", key)
			logger.Info(s)
			return
		}

		for i, k := range data {
			if k.Filename == key {
				data = slices.Delete(data, i, i+1)
				break
			}
		}

		err := saveJSONFile(JSONFILE)
		if err != nil {
			logger.Warn("Error aving data", err)
		}

		s := fmt.Sprintf("Deleting key %s:", key)
		logger.Info(s)
	},
}

var key string

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key to delete")
	deleteCmd.MarkFlagRequired("key")
}
