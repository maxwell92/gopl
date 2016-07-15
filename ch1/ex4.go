package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    var filename string
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, filename = range files {
            counts := make(map[string]int)
            f, err := os.Open(filename)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
            printCounts(counts, filename)
        }
    }
}

func printCounts(counts map[string]int, filename string) {
    for line, n := range counts {
        if n > 1 {
            fmt.Println(filename)
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}