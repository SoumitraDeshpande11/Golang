# Golang Learning Journey 🚀

Welcome to my Go programming adventure! This repository documents my journey learning Go from the ground up.

## 📚 Learning Concepts

The `learning/` folder contains individual Go files, each focusing on a specific concept:

- **00_hello_world.go** - Basic program structure and functions
- **01_variables_and_types.go** - Variables, constants, and data types
- **02_functions.go** - Function declarations, multiple returns, variadic functions
- **03_control_flow.go** - If/else statements, loops, and switch cases
- **04_arrays_slices_maps.go** - Working with arrays, slices, and maps
- **05_structs_methods.go** - Structs and method definitions
- **06_interfaces.go** - Interface implementation and polymorphism
- **07_goroutines.go** - Basic concurrency with goroutines

## 🛠️ Projects

Real-world applications I built while learning:

### Chat Application (`chat-app/`)
A concurrent TCP-based chat server and client supporting multiple users in real-time.
- **Key Concepts**: Goroutines, channels, mutexes, networking
- **Files**: `server.go`, `client.go`

### Markdown to HTML Converter (`markdown-to-html/`)
A CLI tool that converts Markdown files to HTML with support for headers, lists, and inline formatting.
- **Key Concepts**: File I/O, string processing, regular expressions
- **Files**: `main.go`, `sample.md`

## 🎯 Skills Developed

- **Concurrency**: Goroutines, channels, and synchronization
- **Networking**: TCP client-server architecture
- **File Operations**: Reading, writing, and processing files
- **CLI Development**: Building command-line tools
- **Error Handling**: Robust error management patterns
- **Code Organization**: Clean, idiomatic Go code

## 🚀 Getting Started

To run any of the learning examples:
```bash
cd learning/
go run 00_hello_world.go
```

To run the projects:
```bash
# Chat Application
cd chat-app/
go run server.go  # In one terminal
go run client.go  # In another terminal

# Markdown Converter
cd markdown-to-html/
go run main.go sample.md
```

---

*Learning Go one concept at a time! 🐹*
