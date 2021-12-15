package main

import (
	"flag"
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "licenses",
	}
)

func main() {
	flag.Parse()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
