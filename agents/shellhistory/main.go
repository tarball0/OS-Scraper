package shellhistory

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Run() {
	// Set default limit to 10
	n := flag.Int("n", 10, "Number of history items to fetch (default: 10)")
	flag.Parse()
		var HistoryFile string
		HistoryFile = os.Getenv("HOME") + "/.boring_history"

	// If n is <= 0, default to 10
	if *n <= 0 {
		*n = 10
	}

	printLimitedHistory(HistoryFile, *n)
}

func printLimitedHistory(filename string, n int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading shell history file:\n%s\n", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	start := len(lines) - n
	if start < 0 {
		start = 0
	}
	fmt.Printf("Last %d lines of history file:\n", n)
	for _, line := range lines[start:] {
		fmt.Println(line)
	}
}
