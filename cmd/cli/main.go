package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func main() {
	var inputDir, outputDir string
	flag.StringVar(&inputDir, "input", "", "Directory containing Markdown files")
	flag.StringVar(&outputDir, "output", "output", "Output directory for generated HTML files")
	flag.Parse()

	if inputDir == "" {
		fmt.Println("Please specify the directory containing Markdown files using the -input flag.")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		os.Exit(1)
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			os.Exit(1)
		}
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			processMarkdownFile(filepath.Join(inputDir, file.Name()), outputDir)
		}
	}

	fmt.Println("Website generation completed.")
}

func processMarkdownFile(mdFilePath, outputDir string) {
	mdContent, err := ioutil.ReadFile(mdFilePath)
	if err != nil {
		fmt.Printf("Error reading Markdown file %s: %v\n", mdFilePath, err)
		return
	}

	htmlContent := blackfriday.Run(mdContent)

	htmlFileName := strings.TrimSuffix(filepath.Base(mdFilePath), ".md") + ".html"
	htmlFilePath := filepath.Join(outputDir, htmlFileName)

	err = ioutil.WriteFile(htmlFilePath, htmlContent, 0644)
	if err != nil {
		fmt.Printf("Error writing HTML file %s: %v\n", htmlFileName, err)
		return
	}

	fmt.Println("Generated HTML file:", htmlFileName)
}
