package main

import (
	"flag"          // Package flag implements command-line flag parsing.
	"fmt"           // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"io/ioutil"     // Package ioutil implements some I/O utility functions.
	"os"            // Package os provides a platform-independent interface to operating system functionality.
	"path/filepath" // Package filepath implements utility routines for manipulating filename paths.
	"strings"       // Package strings implements simple functions to manipulate UTF-8 encoded strings.

	"github.com/russross/blackfriday/v2" // External package for converting Markdown to HTML.
)

func main() {
	var inputDir, outputDir string

	// Define command-line flags.
	flag.StringVar(&inputDir, "input", "", "Directory containing Markdown files")
	flag.StringVar(&outputDir, "output", "output", "Output directory for generated HTML files")
	flag.Parse() // Parse command-line flags.

	// Check if input directory is specified.
	if inputDir == "" {
		fmt.Println("Please specify the directory containing Markdown files using the -input flag.")
		os.Exit(1)
	}

	// Read the list of files in the input directory.
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		os.Exit(1)
	}

	// Create output directory if it doesn't exist.
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			os.Exit(1)
		}
	}

	// Process each Markdown file in the input directory.
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			processMarkdownFile(filepath.Join(inputDir, file.Name()), outputDir)
		}
	}

	fmt.Println("Website generation completed.")
}

// Process a Markdown file and generate HTML.
func processMarkdownFile(mdFilePath, outputDir string) {
	// Read the content of the Markdown file.
	mdContent, err := ioutil.ReadFile(mdFilePath)
	if err != nil {
		fmt.Printf("Error reading Markdown file %s: %v\n", mdFilePath, err)
		return
	}

	// Convert Markdown to HTML.
	htmlContent := blackfriday.Run(mdContent)

	// Generate HTML file name.
	htmlFileName := strings.TrimSuffix(filepath.Base(mdFilePath), ".md") + ".html"
	// Construct path for HTML file.
	htmlFilePath := filepath.Join(outputDir, htmlFileName)

	// Write HTML content to the HTML file.
	err = ioutil.WriteFile(htmlFilePath, htmlContent, 0644)
	if err != nil {
		fmt.Printf("Error writing HTML file %s: %v\n", htmlFileName, err)
		return
	}

	fmt.Println("Generated HTML file:", htmlFileName)
}
