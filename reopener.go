package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var showStartTime bool
var firstRun = true

func usage() {
	fmt.Println("Usage:")
	s := os.Args[0]
	fmt.Printf("  %s -h       show this memo;\n", s)
	fmt.Printf("  %s [-s] <filename>\n", s)
}

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "show usage memo")
	flag.BoolVar(&showStartTime, "s", false, "show start time")
	flag.Parse()
	if help {
		usage()
		os.Exit(0)
	}
	args := flag.Args()
	if len(args) != 1 {
		usage()
		os.Exit(1)
	}
	period := time.Minute
	if s := os.Getenv("REOPENER_PERIOD"); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("bad period: %s", err)
		}
		period = time.Duration(i) * time.Second
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		openWrite(args[0], period, scanner)
	}
}

func openWrite(path string, period time.Duration, scanner *bufio.Scanner) {
	file, err := os.OpenFile(path, os.O_WRONLY+os.O_APPEND+os.O_CREATE, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if showStartTime && firstRun {
		timestamp := time.Now().Format("2006-01-02T15:04:05.000000-0700")
		_, err := file.Write([]byte(fmt.Sprintf("\n%s *** STARTING ***\n", timestamp)))
		if err != nil {
			log.Fatal(err)
		}
		firstRun = false
	}
	deadline := time.Now().Add(period)
	for time.Now().Before(deadline) {
		if !scanner.Scan() {
			os.Exit(0)
		}
		if _, err := file.Write([]byte(scanner.Text() + "\n")); err != nil {
			log.Fatal(err)
		}
	}
}
