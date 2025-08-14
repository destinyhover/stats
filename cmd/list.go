/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	sort.Sort(DFslice(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		logger.Error("in list()", "err", err)
	}
	fmt.Println(text)
}
func PrettyPrintJSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Implement sort.Interface
func (a DFslice) Len() int {
	return len(a)
}
func (a DFslice) Less(i, j int) bool {
	if a[i].Mean == a[j].Mean {
		return a[i].StdDev < a[j].StdDev
	}
	return a[i].Mean < a[j].Mean
}
func (a DFslice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
