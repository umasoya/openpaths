package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openpaths",
	Short: "A CLI tool to extract and display endpoints from an OpenAPI specification",
	Long: `openpaths is a Go-based CLI application that parses OpenAPI (Swagger) specifications
written in YAML or JSON and lists API endpoints (paths and HTTP methods) in a readable format.

It is useful for API documentation auditing, routing inspection, static analysis,
and generating test cases. Support for resolving external $ref references is also planned.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		fmt.Println(args)
	},
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//func init() {
//	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
