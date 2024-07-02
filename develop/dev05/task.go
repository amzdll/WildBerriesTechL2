package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func grep() {
	var g grepUtil
	g.parseOptions()
	g.openFile()
	g.readFile()
	g.printFile()
}

type grepUtil struct {
	fileName string
	file     *os.File
	pattern  string
	options  options
	content  []string
	data     data
}

type data struct {
	beforeBuff       []string
	afterBuffCounter int
	matchesCount     int
}

type options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func (g *grepUtil) parseOptions() {
	g.fileName = os.Args[len(os.Args)-1]
	if len(os.Args) > 2 {
		g.pattern = os.Args[1]
	}

	// parse options
	flag.IntVar(&g.options.after, "A", 0, "Показывать строки после совпадения")
	flag.IntVar(&g.options.before, "B", 0, "Показывать строки перед совпадением")
	flag.IntVar(&g.options.context, "C", 0, "Показывать строки вокруг совпадения")

	// global parse options
	flag.BoolVar(&g.options.ignoreCase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&g.options.invert, "v", false, "Инвертировать совпадения")
	flag.BoolVar(&g.options.fixed, "F", false, "Точное совпадение строк")
	flag.BoolVar(&g.options.lineNum, "n", false, "Показывать номера строк")

	// output
	flag.BoolVar(&g.options.count, "c", false, "Показывать количество совпадений")

	flag.Parse()
}

func (g *grepUtil) openFile() {
	var err error
	if g.file, err = os.Open(g.fileName); err != nil {
		log.Printf("error opening file: %v", err)
		os.Exit(1)
	}
}

func (g *grepUtil) readFile() {
	line := ""
	lineCount := 1
	bufferLine := make([]string, g.options.before)

	scanner := bufio.NewScanner(g.file)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, g.pattern) {
			if g.options.after > 1 {

			}

			if g.options.before > 0 {
				g.content = append(g.content, bufferLine...)
			}

			if g.options.context > 1 {

			}

			if g.options.ignoreCase {

			}

			if g.options.invert {

			}

			if g.options.fixed {

			}

			if g.options.lineNum {

			}

			if g.options.before > 0 && len(bufferLine) == g.options.before {
				bufferLine = append(bufferLine[1:], line)
			} else {
				bufferLine = append(bufferLine, line)
			}

			g.content = append(g.content, strconv.Itoa(lineCount), line)
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error reading file: %v", err)
		os.Exit(1)
	}
}

func (g *grepUtil) printFile() {
	for _, line := range g.content[:len(g.content)-1] {
		fmt.Println(line)
	}
	fmt.Print(g.content[len(g.content)-1])
}

func main() {
	grep()
}
