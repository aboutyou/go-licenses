package main

import (
	"flag"
	"log"

	"github.com/aboutyou/go-licenses/export"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "licenses",
	}
)

func main() {
	flag.Parse()

	rootCmd.AddCommand(
		export.Command,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
