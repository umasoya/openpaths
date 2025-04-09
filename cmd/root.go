package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
)

// var verbose bool

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
		// parse yaml
		doc := load(args[0])
		// sort paths
		paths := doc.Paths.Map()
		soated := sortPath(doc.Paths)

		// print paths
		for _, path := range soated {
			item := paths[path]
			methods := getAvailableMethods(item)
			for _, method := range methods {
				fmt.Printf("%-6s %s\n", strings.ToUpper(method), path)
			}
		}
	},
	Version: "1.0.0",
}

// load OpenApi specification from file
func load(filepath string) *openapi3.T {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	doc, err := loader.LoadFromFile(filepath)
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}
	if err := doc.Validate(loader.Context); err != nil {
		log.Fatalf("Failed to validate OpenAPI spec: %v", err)
	}
	return doc
}

// sort paths by path
func sortPath(paths *openapi3.Paths) []string {
	var pathList []string
	for path := range paths.Map() {
		pathList = append(pathList, path)
	}
	sort.Strings(pathList)
	return pathList
}

func getAvailableMethods(item *openapi3.PathItem) []string {
	var methods []string
	if item.Get != nil {
		methods = append(methods, "get")
	}
	if item.Post != nil {
		methods = append(methods, "post")
	}
	if item.Put != nil {
		methods = append(methods, "put")
	}
	if item.Patch != nil {
		methods = append(methods, "patch")
	}
	if item.Delete != nil {
		methods = append(methods, "delete")
	}
	if item.Options != nil {
		methods = append(methods, "options")
	}
	if item.Head != nil {
		methods = append(methods, "head")
	}
	if item.Trace != nil {
		methods = append(methods, "trace")
	}
	return methods
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("verbose", "", false, "Enable verbose output")
}
