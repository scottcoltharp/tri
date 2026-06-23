/*
Copyright © 2025 Scott C

*/
package cmd

import (
        "log"
        "fmt"
	"os"

	"github.com/spf13/cobra"
        "github.com/mitchellh/go-homedir"
)

var dataFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
   Use:   "tri",
   Short: "Tri is a todo application",
   Long: `Tri will help you get more done in less
time. It's designed to be as simple as possible to 
help you accomplish your goals.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
         fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

        home, err := homedir.Dir()
        if err != nil {
           log.Println("Unable to detect home directory. Please set data file using --datafile.")
        }
        rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("debug", "d", false, "Add Debug Flag")
}


