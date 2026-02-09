package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type keeper struct {
	timestampFmt    string
	outputFmt       string
	last            time.Time
	n               int
	printTimestamps bool
}

func main() {

	outputFmt := flag.String("o", "%.0f", "floating point output format")
	timestampFmt := flag.String("t", time.RFC3339, "time.Parse timestamp format")
	printTimestamps := flag.Bool("p", false, "print timestamps and intervals")
	flag.Parse()

	fin, closefn, err := openInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closefn()

	realOutputFmt := fmt.Sprintf("%s\n", *outputFmt)

	k := &keeper{
		timestampFmt:    *timestampFmt,
		outputFmt:       realOutputFmt,
		printTimestamps: *printTimestamps,
	}

	scanAllines(fin, k.timestampParser)
}

// scanAllines calls a function (argument fn) on all lines
// of fin argument one at a time. Can print some error messages
// on os.Stderr.
func scanAllines(fin *os.File, fn func(string) error) {

	scanner := bufio.NewScanner(fin)
	/* For longer lines:
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	*/

	lineCounter := 0

	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()
		if err := fn(line); err != nil {
			fmt.Fprintf(os.Stderr, "line %d: %v\n", lineCounter, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "problem line %d: %v", lineCounter, err)
	}
}

// openInputFile either open a file named by os.Args[1],
// and return an *os.File, or if that command line argument doesn't
// exist, return os.Stdin. Also return a closing function, which can
// always be called, even if openInputFile returns os.Stdin
func openInputFile() (*os.File, func(), error) {
	var closeFunc = func() {}
	fin := os.Stdin
	if flag.NArg() > 0 {
		var err error
		if fin, err = os.Open(flag.Arg(0)); err != nil {
			return nil, closeFunc, err
		}
		closeFunc = func() { fin.Close() }
	}
	return fin, closeFunc, nil
}

func (k *keeper) timestampParser(text string) error {
	// fmt.Printf("# %s\n", text)
	// timestamp, err := time.Parse("2006-01-02T15:04:05-07:00", text)
	timestamp, err := time.Parse(k.timestampFmt, text)
	if err != nil {
		return fmt.Errorf("%q unparseable\n", text)
	}
	if k.n > 0 {
		x := timestamp.Sub(k.last)
		fmt.Printf(k.outputFmt, x.Seconds())
	}
	if k.printTimestamps {
		fmt.Printf("%s\n", text)
	}
	k.last = timestamp
	k.n++
	return nil
}
