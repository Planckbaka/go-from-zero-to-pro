package main

import (
	"fmt"
	"os"
)

// CLI for Go from Zero to Pro learning tool.
func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "list":
		listLessons()
	case "run":
		runLesson()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`z2p - Go from Zero to Pro CLI

Usage:
  z2p <command> [arguments]

Commands:
  list      List all available lessons
  run       Run a specific lesson
  help      Show this help message

Examples:
  z2p list
  z2p run basics`)
}

func listLessons() {
	lessons := []struct {
		id   string
		name string
		desc string
	}{
		{"00", "Introduction", "Get started with Go"},
		{"01", "Basics", "Learn the fundamentals"},
		{"02", "Types", "Understand Go's type system"},
		{"03", "Functions", "Master functions and methods"},
		{"04", "Interfaces", "Explore interfaces and polymorphism"},
		{"05", "Concurrency", "Dive into concurrent programming"},
		{"90", "Tooling", "Development tools and practices"},
		{"91", "Testing", "Write effective tests"},
		{"92", "Performance", "Optimize your code"},
		{"93", "Security", "Build secure applications"},
	}

	fmt.Println("Available Lessons:")
	fmt.Println("==================")
	for _, l := range lessons {
		fmt.Printf("  [%s] %s - %s\n", l.id, l.name, l.desc)
	}
}

func runLesson() {
	fmt.Println("Run lesson: coming soon!")
	fmt.Println("For now, explore the docs/ and examples/ directories.")
}
