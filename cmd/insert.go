/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insertCmd",
	Long:  `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		logger = GetLogger()

		if file == "" {
			logger.Info("Need a file to read!")
			return
		}

		_, ok := index[file]
		if ok {
			for i, k := range data {
				if k.Filename == file {
					data = slices.Delete(data, i, i+1)
					break
				}
			}
		}

		err := ProcessFile(file)
		if err != nil {
			logger.Error("Error processing:", "err:", err)
		}

		err = saveJSONFile(JSONFILE)
		if err != nil {
			logger.Info("Error saving data:", "err:", err)
		}

	},
}
var file string

func init() {
	logger := GetLogger()
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringVarP(&file, "file", "f", "", "Filename to process")
	insertCmd.MarkFlagRequired("file")
	s := fmt.Sprintf("%d records in total.", len(data))
	logger.Info(s)
}

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			logger.Error("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}
	return values, nil
}

func stDev(x []float64) (float64, float64) {
	sum := float64(0)
	for _, val := range x {
		sum = sum + val
	}
	meanValue := sum / float64(len(x))

	// Stadndard deviation
	var squared float64
	for i := 0; i < len(x); i++ {
		squared = squared + math.Pow((x[i]-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(len(x)))
	return meanValue, standardDeviation
}

func ProcessFile(file string) error {
	currentFile := Entry{}
	currentFile.Filename = file

	values, err := readFile(file)
	if err != nil {
		return err
	}

	currentFile.Len = len(values)
	currentFile.Minimum = slices.Min(values)
	currentFile.Maximum = slices.Max(values)
	meanValue, standardDeviation := stDev(values)
	currentFile.Mean = meanValue
	currentFile.StdDev = standardDeviation

	data = append(data, currentFile)
	return nil
}
