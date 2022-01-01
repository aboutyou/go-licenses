package main

import (
	"flag"
	"log"

	"github.com/aboutyou/go-licenses/check"
	"github.com/aboutyou/go-licenses/export"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-licenses",
		Short: "go-licenses is a license compliance tool for Go Modules",
	}
)

func main() {
	flag.Parse()

	rootCmd.AddCommand(
		export.Command,
		check.Command,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
