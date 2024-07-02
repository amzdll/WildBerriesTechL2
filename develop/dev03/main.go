package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

type sorter struct {
	fileName string
	content  []string
	options  options
}

type options struct {
	k int
	n bool
	r bool
	u bool
}

func Sort() {
	var s sorter

	s.parseArgs()
	s.readFile()

	if s.options.u {
		s.unique()
	}

	if s.options.k > 1 {
		sort.Slice(s.content, func(i, j int) bool {
			lCmp := strings.Split(s.content[i], " ")
			rCmp := strings.Split(s.content[j], " ")
			if s.options.k >= len(lCmp) {
				return true
			}
			if s.options.k < len(rCmp) {
				return lCmp[s.options.k] < rCmp[s.options.k]
			}
			return false
		})
	}

	if s.options.r {
		slices.Reverse(s.content)
	}

	fmt.Println(slices.IsSorted(s.content))

	s.printFile()
}

func (s *sorter) unique() {
	var result []string

	slices.Sort(s.content)

	for _, str := range s.content {
		i := sort.Search(len(result), func(i int) bool {
			return result[i] >= str
		})

		if i == len(result) {
			result = append(result, str)
		} else {
			if result[i] != str {
				result = append(result[:i], append([]string{str}, result[i:]...)...)
			}
		}
	}

	s.content = result
}

func (s *sorter) parseArgs() {
	s.fileName = os.Args[len(os.Args)-1]

	flag.IntVar(&s.options.k, "k", 1, "column number for sorting")
	flag.BoolVar(&s.options.n, "n", false, "sort numerically")
	flag.BoolVar(&s.options.r, "r", false, "sort in reverse order")
	flag.BoolVar(&s.options.u, "u", false, "output unique lines only")
	flag.Parse()

	s.options.k -= 1
}

func (s *sorter) readFile() {
	file, err := os.Open(s.fileName)
	if err != nil {
		log.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s.content = append(s.content, strings.TrimSpace(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		log.Printf("error reading file: %v", err)
		os.Exit(1)
	}

}

func (s *sorter) printFile() {
	for _, line := range s.content[:len(s.content)-1] {
		fmt.Println(line)
	}
	fmt.Print(s.content[len(s.content)-1])
}

func main() {
	Sort()
}
