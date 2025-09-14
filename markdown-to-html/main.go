package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <markdown-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := strings.TrimSuffix(inputFile, ".md") + ".html"

	content, err := readFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	html := convertToHTML(content)
	
	err = writeFile(outputFile, html)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Converted %s to %s\n", inputFile, outputFile)
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	return content.String(), scanner.Err()
}

func writeFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func convertToHTML(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var html strings.Builder
	
	html.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<title>Converted Document</title>\n</head>\n<body>\n")
	
	inCodeBlock := false
	inList := false
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				html.WriteString("</code></pre>\n")
				inCodeBlock = false
			} else {
				html.WriteString("<pre><code>\n")
				inCodeBlock = true
			}
			continue
		}
		
		if inCodeBlock {
			html.WriteString(line + "\n")
			continue
		}
		
		if line == "" {
			if inList {
				html.WriteString("</ul>\n")
				inList = false
			}
			html.WriteString("<br>\n")
			continue
		}
		
		if strings.HasPrefix(line, "# ") {
			if inList {
				html.WriteString("</ul>\n")
				inList = false
			}
			html.WriteString(fmt.Sprintf("<h1>%s</h1>\n", line[2:]))
		} else if strings.HasPrefix(line, "## ") {
			if inList {
				html.WriteString("</ul>\n")
				inList = false
			}
			html.WriteString(fmt.Sprintf("<h2>%s</h2>\n", line[3:]))
		} else if strings.HasPrefix(line, "### ") {
			if inList {
				html.WriteString("</ul>\n")
				inList = false
			}
			html.WriteString(fmt.Sprintf("<h3>%s</h3>\n", line[4:]))
		} else if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			if !inList {
				html.WriteString("<ul>\n")
				inList = true
			}
			html.WriteString(fmt.Sprintf("<li>%s</li>\n", processInlineFormatting(line[2:])))
		} else {
			if inList {
				html.WriteString("</ul>\n")
				inList = false
			}
			html.WriteString(fmt.Sprintf("<p>%s</p>\n", processInlineFormatting(line)))
		}
	}
	
	if inList {
		html.WriteString("</ul>\n")
	}
	
	html.WriteString("</body>\n</html>")
	return html.String()
}

func processInlineFormatting(text string) string {
	boldRegex := regexp.MustCompile(`\*\*(.*?)\*\*`)
	text = boldRegex.ReplaceAllString(text, "<strong>$1</strong>")
	
	italicRegex := regexp.MustCompile(`\*(.*?)\*`)
	text = italicRegex.ReplaceAllString(text, "<em>$1</em>")
	
	codeRegex := regexp.MustCompile("`(.*?)`")
	text = codeRegex.ReplaceAllString(text, "<code>$1</code>")
	
	linkRegex := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	text = linkRegex.ReplaceAllString(text, `<a href="$2">$1</a>`)
	
	return text
}
