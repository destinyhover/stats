/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type Entry struct {
	Filename string  `json:"filename"`
	Len      int     `json:"length"`
	Minimum  float64 `json:"minimum"`
	Maximum  float64 `json:"maximum"`
	Mean     float64 `json:"mean"`
	StdDev   float64 `json:"stddev"`
}

var logger *slog.Logger

var JSONFILE = "./data.json"

type DFslice []Entry

var data = DFslice{}
var index map[string]int

// ----------- JSON FUNCS
func DeSerialize(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}
func Serialize(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = Serialize(&data, f)
	return err

}
func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	return nil
}
func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Filename
		index[key] = i
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your application",
	Long:  `A longer description `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// l := slog.Default().With("cmd", cmd.CommandPath(), "pid", os.Getpid())
		// ctx := slog.NewContext(cmd.Context(), l)
		// cmd.SetContext(ctx)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
func initLogger() {
	out := io.Writer(os.Stderr)
	// Work with logger
	if !enableLogging {
		out = io.Discard
	}
	logger = slog.New(slog.NewJSONHandler(out, nil))
	slog.SetDefault(logger)

}

//	func logFromCmd(cmd *cobra.Command) *slog.Logger {
//		// if l := slog.FromContext(cmd.Context()); l != nil {
//		// 	return l
//		// }
//		return slog.Default()
//	}
func getLogger() *slog.Logger {
	if logger != nil {
		return logger
	}
	return slog.Default()
}

var enableLogging bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&enableLogging, "log", "l", true, "Logging information")
	cobra.OnInitialize(initLogger)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	err := readJSONFile(JSONFILE)
	if err != nil && err != io.EOF {
		return
	}
	createIndex()
}
