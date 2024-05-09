package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

type Args struct {
	Test string
	Debug bool
}


func main() {

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true,
		ReportTimestamp: true,
		TimeFormat: time.Kitchen,
		Prefix: "CookieLove 🍪 ",
	})


	// set arguments
	var arguments Args

	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "A brief description of your application",
		Long:  "A longer description that spans multiple lines and likely contains examples and usage of using your application. For example, `myapp scan` or `myapp analyze`.",
		Run: func(cmd *cobra.Command, args []string) {
			if arguments.Debug {
				logger.SetLevel(log.DebugLevel)
			}
			logger.Info("Starting oven!", "degree", 375)
		},
	}
	rootCmd.Flags().StringVarP(&arguments.Test, "test", "p", "", "test love")
	rootCmd.Flags().BoolVarP(&arguments.Debug, "debug", "d", false, "Start program with debug")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
