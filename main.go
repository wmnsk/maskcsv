package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/csv"
	"flag"
)

var (
	infile   = flag.String("i", "", "File to be masked.")
	indelim  = flag.String("d", ",", "Delimiter of input file.")
	outfile  = flag.String("o", "", "File to be masked.")
	outdelim = flag.String("s", ",", "Delimiter of output file.")
	masklen  = flag.Int("l", 2, "Number of the letters to mask.")
	maskchar = flag.String("m", "X", "Character to be used as the mask.")
	fields   = flag.String("f", "no,header,val", "Path to the original CSV file to be masked.")
)

func readInputFile(f string, del rune) [][]string {
	fp, err := os.Open(f)
	if err != nil {
		fmt.Println("ERROR: No input file specified :-(")
		panic(err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	r.Comma = del
	r.LazyQuotes = true

	rcds, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	return rcds
}

func writeAsCSV(f string, del rune, masked [][]string) int {
	fp, err := os.Create(f)
	if err != nil {
		fmt.Println("ERROR: No output filename specified :-(")
		panic(err)
	}
	defer fp.Close()

	w := csv.NewWriter(fp)
	w.Comma = del

	w.WriteAll(masked)
	return 0
}

func getFieldIndex(flds []string, hdr []string) map[string]int {
	hdrs := make(map[string]int)
	for _, f := range flds {
		for n, r := range hdr {
			if f == r {
				hdrs[r] = n
			}
		}
	}
	return hdrs
}

func maskLastLetters(s string, m string, l int) string {
	rem := s[0 : len(s)-l]
	msk := strings.Repeat(m, l)
	return rem + msk
}

func main() {
	flag.Parse()
	indel := []rune(*indelim)[0]
	outdel := []rune(*outdelim)[0]
	fields := strings.Split(*fields, ",")

	records := readInputFile(*infile, indel)
	indice := getFieldIndex(fields, records[0])

	masked := make([][]string, len(records))
	for i, line := range records {
		for _, v := range indice {
			if (i == 0) || (len(line[v]) == 0) || (len(line[v]) < *masklen) {
			} else {
				line[v] = maskLastLetters(line[v], *maskchar, *masklen)
			}
		}
		masked[i] = line
	}

	result := writeAsCSV(*outfile, outdel, masked)
	if result == 0 {
		fmt.Println("Masking successfully finished!")
	} else {
		fmt.Println("Something wrong happened...")
	}
}
