package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Playground struct {
	tempDir string
}

func NewPlayground() *Playground {
	tempDir := filepath.Join(os.TempDir(), "go-playground")
	os.MkdirAll(tempDir, 0755)
	return &Playground{tempDir: tempDir}
}

func (p *Playground) RunCode(code string) (string, error) {
	filename := filepath.Join(p.tempDir, fmt.Sprintf("main_%d.go", time.Now().UnixNano()))
	
	err := os.WriteFile(filename, []byte(code), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write code to file: %v", err)
	}
	defer os.Remove(filename)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", filename)
	cmd.Dir = p.tempDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("code execution timed out")
		}
		return string(output), fmt.Errorf("execution failed: %v", err)
	}

	return string(output), nil
}

func (p *Playground) ValidateCode(code string) error {
	if !strings.Contains(code, "package main") {
		return fmt.Errorf("code must contain 'package main'")
	}
	
	if !strings.Contains(code, "func main()") {
		return fmt.Errorf("code must contain 'func main()'")
	}

	dangerousPatterns := []string{
		"os.Remove",
		"os.RemoveAll",
		"exec.Command",
		"syscall",
		"unsafe",
		"net/http",
		"net.Listen",
		"os.Exit",
	}

	for _, pattern := range dangerousPatterns {
		if strings.Contains(code, pattern) {
			return fmt.Errorf("code contains potentially dangerous operation: %s", pattern)
		}
	}

	return nil
}

func main() {
	playground := NewPlayground()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Go Playground Clone - Enter your Go code (type 'RUN' on a new line to execute, 'QUIT' to exit)")
	fmt.Println("Note: Only safe operations are allowed. No file system, network, or system calls.")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("Enter Go code:\n")
		
		var codeLines []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "RUN" {
				break
			}
			if line == "QUIT" {
				fmt.Println("Goodbye!")
				return
			}
			if line == "CLEAR" {
				codeLines = []string{}
				fmt.Println("Code cleared. Enter new code:")
				continue
			}
			if line == "HELP" {
				printHelp()
				continue
			}
			codeLines = append(codeLines, line)
		}

		if len(codeLines) == 0 {
			fmt.Println("No code entered. Try again.")
			continue
		}

		code := strings.Join(codeLines, "\n")
		
		if err := playground.ValidateCode(code); err != nil {
			fmt.Printf("Validation Error: %v\n", err)
			fmt.Println("----------------------------------------")
			continue
		}

		fmt.Println("Running code...")
		fmt.Println("Output:")
		fmt.Println("========================================")
		
		output, err := playground.RunCode(code)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			if output != "" {
				fmt.Printf("Output: %s\n", output)
			}
		} else {
			if output == "" {
				fmt.Println("(no output)")
			} else {
				fmt.Print(output)
			}
		}
		
		fmt.Println("========================================")
		fmt.Println()
	}
}

func printHelp() {
	fmt.Println(`
Available commands:
- RUN: Execute the entered code
- CLEAR: Clear the current code buffer
- QUIT: Exit the playground
- HELP: Show this help message

Example code:
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
RUN
`)
}
